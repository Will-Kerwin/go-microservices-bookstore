package author

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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

const GetAuthorsBaseKey string = "Authors"

// HTTP Handler for author endpoints
type Handler struct {
	gateway  gateway.AuthorGateway
	kafkaUri string
	redis    *redis.Client
}

func (h *Handler) invalidateAuthorCache(ctx context.Context) {
	keys, err := h.redis.Keys(ctx, GetAuthorsBaseKey+"*").Result()
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

// Create a new instance of the handler
func New(gateway gateway.AuthorGateway, redis *redis.Client, kafkaUri string) *Handler {
	return &Handler{
		gateway:  gateway,
		kafkaUri: kafkaUri,
		redis:    redis,
	}
}

func (h *Handler) newProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": h.kafkaUri})
}

// Register endpoints for the handler
func (h *Handler) Register(r *echo.Group) {
	r.GET("/authors", h.GetAuthors)
	r.GET("/authors/:id", h.GetAuthor)
	r.POST("/authors", h.CreateAuthor)
	r.DELETE("/authors/:id", h.DeleteAuthor)
}

// GetAuthors godoc
// @Summary Get Authors.
// @Description get the authors from database.
// @Tags authors
// @Accept applicaiton/json
// @Produce json
// @Success 200 {object} []models.Author
// @Failure 502 {object} models.ApiErrorResponse
// @Router /authors [get]
func (h *Handler) GetAuthors(ctx echo.Context) error {

	val, err := h.redis.Get(ctx.Request().Context(), GetAuthorsBaseKey).Result()

	if err != nil {
		log.Println("Authors Cache Miss")
		log.Println(err)
	}

	if val != "" {
		var authors []models.Author

		if err = json.Unmarshal([]byte(val), &authors); err == nil {
			log.Println("Authors cache hit")
			return ctx.JSON(http.StatusOK, authors)
		} else {
			log.Println(err)
		}
	}

	res, err := h.gateway.Get(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	marshaledRes, err := json.Marshal(res)

	if err == nil {
		log.Printf("update authors cache")
		h.redis.Set(ctx.Request().Context(), GetAuthorsBaseKey, marshaledRes, time.Minute*10).Err()
	}

	return ctx.JSON(http.StatusOK, res)
}

// GetAuthor godoc
// @Summary Get Author by its object id in hex format.
// @Description get the author by id from database.
// @Tags authors
// @Accept applicaiton/json
// @Produce json
// @Param  id path string true "id of the author"
// @Success 200 {object} models.Author
// @Failure 400 {object} models.ApiErrorResponse
// @Failure 404 {object} models.ApiErrorResponse
// @Failure 502 {object} models.ApiErrorResponse
// @Router /authors/{id} [get]
func (h *Handler) GetAuthor(ctx echo.Context) error {

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

// CreateAuthor godoc
// @Summary Create an author.
// @Description creates an author asynchronously.
// @Tags authors
// @Accept applicaiton/json
// @Produce json
// @Param  body body models.Author true "author body"
// @Success 201
// @Success 400 {object} models.ApiErrorResponse
// @Success 502 {object} models.ApiErrorResponse
// @Router /authors [post]
func (h *Handler) CreateAuthor(ctx echo.Context) error {
	topicName := "createAuthor"
	producer, err := h.newProducer()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}
	defer producer.Close()

	author := new(models.Author)

	if err := ctx.Bind(author); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "could not parse body"})
	}

	encodedEvent, err := json.Marshal(events.CreateAuthorEvent{
		Name:        author.Name,
		DateOfBirth: author.DateOfBirth,
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
	h.invalidateAuthorCache(ctx.Request().Context())

	return ctx.NoContent(http.StatusAccepted)

}

// DeleteAuthor godoc
// @Summary Delete Author by its object id in hex format.
// @Description delete the author by id from database.
// @Tags authors
// @Accept applicaiton/json
// @Produce json
// @Param  id path string true "id of the author"
// @Success 202
// @Failure 400 {object} models.ApiErrorResponse
// @Failure 502 {object} models.ApiErrorResponse
// @Router /authors/{id} [delete]
func (h *Handler) DeleteAuthor(ctx echo.Context) error {
	id := ctx.Param("id")

	topicName := "deleteAuthor"
	producer, err := h.newProducer()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}
	defer producer.Close()

	encodedEvent, err := json.Marshal(events.DeleteAuthorEvent{ID: id})
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
	h.invalidateAuthorCache(ctx.Request().Context())

	return ctx.NoContent(http.StatusAccepted)
}
