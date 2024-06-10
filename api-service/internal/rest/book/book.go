package book

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HTTP Handler for book endpoints
type Handler struct {
	gateway  gateway.BookGateway
	kafkaUri string
}

// Create a new instance of the handler
func New(gateway gateway.BookGateway, kafkaUri string) *Handler {
	return &Handler{gateway: gateway, kafkaUri: kafkaUri}
}

// Register book endpoints
func (h *Handler) Register(r *echo.Echo) {
	r.GET("/books", h.GetBooks)
	r.GET("/books/:id", h.GetBook)
	r.POST("/books", h.CreateBook)
	r.PATCH("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)
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

	res, err := h.gateway.Get(ctx.Request().Context(), title, authorId, genre)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
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
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": h.kafkaUri})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}
	defer producer.Close()

	book := new(models.Book)

	if err := ctx.Bind(book); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "could not parse body"})
	}

	encodedEvent, err := json.Marshal(models.CreateBookEvent{
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

	producer.Flush(int((1 * time.Second).Milliseconds()))

	return ctx.NoContent(http.StatusCreated)
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
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": h.kafkaUri})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}
	defer producer.Close()

	book := new(models.Book)

	if err := ctx.Bind(book); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "could not parse body"})
	}

	encodedEvent, err := json.Marshal(models.UpdateBookEvent{
		Data: models.UpdateBookEventData{
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
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": h.kafkaUri})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}
	defer producer.Close()

	encodedEvent, err := json.Marshal(models.DeleteBookEvent{ID: id})
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

	return ctx.NoContent(http.StatusNoContent)
}
