package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTP Handler for book endpoints
type Handler struct {
	// poss controller
}

// Create a new instance of the handler
func New() *Handler {
	return &Handler{}
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
	return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"test": "tests"})
}

// Get a book by its id
func (h *Handler) GetBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusBadRequest)
}

// Create a book
func (h *Handler) CreateBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusBadRequest)
}

// Update a books information
func (h *Handler) UpdateBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusBadRequest)
}

// delete a book
func (h *Handler) DeleteBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusBadRequest)
}
