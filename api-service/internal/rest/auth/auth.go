package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/rest/middleware"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gateway  gateway.AuthGateway
	kafkaUri string
}

func New(authGateway gateway.AuthGateway, kafkaUri string) *Handler {
	return &Handler{
		gateway:  authGateway,
		kafkaUri: kafkaUri,
	}
}

func (h *Handler) Register(r *echo.Echo) {
	r.POST("/auth/login", h.Login)
	r.POST("/auth/users", h.CreateUser)

	userSpecificGroup := r.Group("/auth/users/:id")

	userSpecificGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return middleware.UseAdminOrSameUserAuthMiddleware
	})

	userSpecificGroup.GET("/auth/users/:id", h.GetUser)
	userSpecificGroup.PATCH("/auth/users/:id", h.UpdateUser)
}

func (h *Handler) newProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": h.kafkaUri})
}

// Login godoc
// @Summary Login
// @Description login to the api
// @Tags auth
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} models.LoginResponse
// @Failure 401 {object} models.ApiErrorResponse
// @Router /auth/login [post]
func (h *Handler) Login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	user, err := h.gateway.LoginUser(ctx.Request().Context(), username, password)

	if err != nil {
		switch status.Code(err) {
		case codes.Unauthenticated:
			return ctx.JSON(http.StatusUnauthorized, models.ApiErrorResponse{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
		}
	}

	jwt, err := models.BuildJwt(user.Username, user.Email)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": jwt,
	})
}

// GetUser godoc
// @Summary Get user by its object id in hex format.
// @Description get the user by id from database.
// @Tags auth
// @Accept applicaiton/json
// @Produce json
// @Param  id path string true "id of the user"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ApiErrorResponse
// @Failure 404 {object} models.ApiErrorResponse
// @Failure 502 {object} models.ApiErrorResponse
// @Router /auth/users/{id} [get]
func (h *Handler) GetUser(ctx echo.Context) error {
	id := ctx.Param("id")

	res, err := h.gateway.GetUser(ctx.Request().Context(), id)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return ctx.JSON(http.StatusNotFound, models.ApiErrorResponse{"error": err.Error()})
			default:
				log.Printf("Get User: failed: Err: %v\n", err)
				return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": err.Error()})
			}
		}

		log.Printf("not able to parse error returned %v", err)

		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

// CreateUser godoc
// @Summary CreateUser
// @Description create a new user
// @Tags auth
// @Accept application/json
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} models.ApiErrorResponse
// @Router /auth/users [post]
func (h *Handler) CreateUser(ctx echo.Context) error {
	topicName := "createUser"

	producer, err := h.newProducer()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	defer producer.Close()
	createReq := new(models.CreateUserEvent)

	if err := ctx.Bind(createReq); err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "could not parse body"})
	}

	if createReq.Username == "" {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "email is invalid"})
	}

	isValid, err := h.gateway.ValidateUsernameUnique(ctx.Request().Context(), createReq.Username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ApiErrorResponse{"error": err.Error()})
	}

	if !isValid {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "username already exists"})
	}

	if createReq.Email == "" || !models.EmailRegex.MatchString(createReq.Email) {
		return ctx.JSON(http.StatusBadRequest, models.ApiErrorResponse{"error": "email is invalid"})
	}

	encodedEvent, err := json.Marshal(createReq)

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

// UpdateUser godoc
// @Summary UpdateUser
// @Description Update an existing user
// @Tags auth
// @Accept application/json
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} models.ApiErrorResponse
// @Router /auth/users/:id [patch]
func (h *Handler) UpdateUser(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
