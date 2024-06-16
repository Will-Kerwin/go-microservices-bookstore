package book

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/events"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const GetBooksBaseKey string = "Books"

// HTTP Handler for book endpoints
type Handler struct {
	gateway  gateway.BookGateway
	kafkaUri string
	redis    *redis.Client
}

// Create a new instance of the handler
func New(gateway gateway.BookGateway, redist *redis.Client, kafkaUri string) *Handler {
	return &Handler{gateway: gateway, kafkaUri: kafkaUri, redis: redist}
}

func (h *Handler) newProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": h.kafkaUri})
}

// Register book endpoints
func (h *Handler) Register(r *echo.Group) {
	r.GET("/books", h.GetBooks)
	r.GET("/books/:id", h.GetBook)
	r.POST("/books", h.CreateBook)
	r.PATCH("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)
}

// getBooksKey godoc
// Get the cache key for books
func getBooksKey(title string, authorId string, genre string) string {
	key := GetBooksBaseKey

	if genre != "" {
		key += fmt.Sprintf(":%s", genre)
	} else {
		key += ":"
	}

	if title != "" {
		key += fmt.Sprintf(":%s", title)
	} else {
		key += ":"
	}

	if authorId != "" {
		key += fmt.Sprintf(":%s", authorId)
	} else {
		key += ":"
	}

	key = strings.Trim(key, ":")

	h := sha256.New()

	h.Write([]byte(key))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (h *Handler) invalidateBookCache(ctx context.Context) {
	keys, err := h.redis.Keys(ctx, GetBooksBaseKey+"*").Result()
	if err != nil {
		return
	}

	if len(keys) == 0 {
		return
	}

	err = h.redis.Del(ctx, keys...).Err()

	if err != nil {
		log.Println(err)
	}
}

// GetBooks godoc
// @Summary Get Books.
// @Description get the books from database with filters.
// @Tags books
// @Accept applicaiton/json
// @Param  title path string false "title of the book"
// @Param  genre path string false "genre of the book"
// @Param  authorId path string false "authorId of the book"
// @Produce json
// @Success 200 {object} []models.Book
// @Failure 502 {object} models.ApiErrorResponse
// @Router /books [get]
func (h *Handler) GetBooks(ctx echo.Context) error {

	genre := ctx.Param("genre")
	title := ctx.Param("title")
	authorId := ctx.Param("authorId")

	key := getBooksKey(title, authorId, genre)

	val, err := h.redis.Get(ctx.Request().Context(), key).Result()

	if err != nil {
		log.Println("Books Cache Miss")
		log.Println(err)
	}

	if val != "" {
		var books []models.Book

		if err = json.Unmarshal([]byte(val), &books); err == nil {
			log.Println("Books cache hit")
			return ctx.JSON(http.StatusOK, books)
		} else {
			log.Println(err)
		}
	}

	res, err := h.gateway.Get(ctx.Request().Context(), title, authorId, genre)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	marshaledRes, err := json.Marshal(res)

	if err == nil {
		log.Printf("update authors cache")
		h.redis.Set(ctx.Request().Context(), key, marshaledRes, time.Minute*10).Err()
	}

	return ctx.JSON(http.StatusOK, res)
}

// GetBook godoc
// @Summary Get book by its object id in hex format.
// @Description get the book by id from database.
// @Tags books
// @Accept applicaiton/json
// @Produce json
// @Param  id path string true "id of the book"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ApiErrorResponse
// @Failure 404 {object} models.ApiErrorResponse
// @Failure 502 {object} models.ApiErrorResponse
// @Router /books/{id} [get]
func (h *Handler) GetBook(ctx echo.Context) error {
	id := ctx.Param("id")

	res, err := h.gateway.GetById(ctx.Request().Context(), id)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return ctx.JSON(http.StatusNotFound, models.ApiErrorResponse{"error": err.Error()})
			default:
				log.Printf("GetAuthor failed: Err: %v\n", err)
				return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": err.Error()})
			}
		}

		log.Printf("not able to parse error returned %v", err)

		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

// CreateBook godoc
// @Summary Create an book.
// @Description creates a book asynchronously.
// @Tags books
// @Accept applicaiton/json
// @Produce json
// @Param  body body models.Book true "book body"
// @Success 201
// @Success 400 {object} models.ApiErrorResponse
// @Success 502 {object} models.ApiErrorResponse
// @Router /books [post]
func (h *Handler) CreateBook(ctx echo.Context) error {

	topicName := "createBook"
	producer, err := h.newProducer()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}
	defer producer.Close()

	book := new(models.Book)

	if err := ctx.Bind(book); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "could not parse body"})
	}

	encodedEvent, err := json.Marshal(events.CreateBookEvent{
		Title:    book.Title,
		AuthorId: book.AuthorId,
		Synopsis: book.Synopsis,
		ImageUrl: book.ImageUrl,
		Genre:    book.Genre,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": err.Error()})
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Value:          encodedEvent,
	}

	if err := producer.Produce(message, nil); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	h.invalidateBookCache(ctx.Request().Context())
	producer.Flush(int((1 * time.Second).Milliseconds()))

	return ctx.NoContent(http.StatusAccepted)
}

// UpdateBook godoc
// @Summary Update book by its object id in hex format.
// @Description Update the book by id from database.
// @Tags books
// @Accept applicaiton/json
// @Produce json
// @Param  id path string true "id of the book"
// @Param  body body models.Book true "body of the book"
// @Success 202
// @Failure 400 {object} models.ApiErrorResponse
// @Failure 502 {object} models.ApiErrorResponse
// @Router /books/{id} [patch]
func (h *Handler) UpdateBook(ctx echo.Context) error {
	id := ctx.Param("id")
	topicName := "updateBook"
	producer, err := h.newProducer()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}
	defer producer.Close()

	book := new(models.Book)

	if err := ctx.Bind(book); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "could not parse body"})
	}

	encodedEvent, err := json.Marshal(events.UpdateBookEvent{
		Data: events.UpdateBookEventData{
			Title:    book.Title,
			AuthorId: book.AuthorId,
			Synopsis: book.Synopsis,
			ImageUrl: book.ImageUrl,
			Genre:    book.Genre,
		},
		ID: id,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": err.Error()})
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Value:          encodedEvent,
	}

	if err := producer.Produce(message, nil); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	producer.Flush(int((1 * time.Second).Milliseconds()))
	h.invalidateBookCache(ctx.Request().Context())
	return ctx.NoContent(http.StatusAccepted)

}

// DeleteBook godoc
// @Summary Delete book by its object id in hex format.
// @Description delete the book by id from database.
// @Tags books
// @Accept applicaiton/json
// @Produce json
// @Param  id path string true "id of the book"
// @Success 202
// @Failure 400 {object} models.ApiErrorResponse
// @Failure 502 {object} models.ApiErrorResponse
// @Router /books/{id} [delete]
func (h *Handler) DeleteBook(ctx echo.Context) error {
	id := ctx.Param("id")

	topicName := "deleteBook"
	producer, err := h.newProducer()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}
	defer producer.Close()

	encodedEvent, err := json.Marshal(events.DeleteBookEvent{ID: id})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": err.Error()})
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Value:          encodedEvent,
	}

	if err := producer.Produce(message, nil); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	producer.Flush(int((1 * time.Second).Milliseconds()))
	h.invalidateBookCache(ctx.Request().Context())
	return ctx.NoContent(http.StatusAccepted)
}
