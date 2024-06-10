package book

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HTTP Handler for book endpoints
type Handler struct {
	gateway gateway.BookGateway
}

// Create a new instance of the handler
func New(gateway gateway.BookGateway) *Handler {
	return &Handler{gateway: gateway}
}

// Register book endpoints
func (h *Handler) Register(r *echo.Echo) {
	r.GET("/books", h.GetBooks)
	r.GET("/books/:id", h.GetBook)
	r.POST("/books", h.CreateBook)
	r.PATCH("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)
}

// Get books
func (h *Handler) GetBooks(ctx echo.Context) error {

	genre := ctx.Param("genre")
	title := ctx.Param("title")
	authorId := ctx.Param("authorId")

	res, err := h.gateway.Get(ctx.Request().Context(), title, authorId, genre)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

// Get a book by its id
func (h *Handler) GetBook(ctx echo.Context) error {
	id := ctx.Param("id")

	res, err := h.gateway.GetById(ctx.Request().Context(), id)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			default:
				log.Printf("GetAuthor failed: Err: %v\n", err)
				return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}
		}

		log.Printf("not able to parse error returned %v", err)

		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

// Create a book
func (h *Handler) CreateBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)
}

// Update a books information
func (h *Handler) UpdateBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)
}

// delete a book
func (h *Handler) DeleteBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
