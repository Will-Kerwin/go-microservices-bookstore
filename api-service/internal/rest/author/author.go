package author

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/will-kerwin/go-microservice-bookstore/api-service/internal/gateway"

	// "github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HTTP Handler for author endpoints
type Handler struct {
	gateway gateway.AuthorGateway
}

// Create a new instance of the handler
func New(gateway gateway.AuthorGateway) *Handler {
	return &Handler{
		gateway: gateway,
	}
}

// Register endpoints for the handler
func (h *Handler) Register(r *echo.Echo) {
	r.GET("/authors", h.GetAuthors)
	r.GET("/authors/:id", h.GetAuthor)
	r.POST("/authors", h.CreateAuthor)
	r.DELETE("/authors/:id", h.DeleteAuthor)
}

// Get all authors
func (h *Handler) GetAuthors(ctx echo.Context) error {
	res, err := h.gateway.Get(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

// Get an author by its ID
func (h *Handler) GetAuthor(ctx echo.Context) error {

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

// Create a new author
func (h *Handler) CreateAuthor(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)

	// author := new(models.Author)

	// if err := ctx.Bind(author); err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "could not parse body"})
	// }

	// author, err := h.gateway.CreateAuthor(ctx.Request().Context(), author)

	// if err != nil {
	// 	if e, ok := status.FromError(err); ok {
	// 		switch e.Code() {
	// 		case codes.NotFound:
	// 			return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
	// 		default:
	// 			log.Printf("CreateAuthor failed: Err: %v\n", err)
	// 			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	// 		}
	// 	}

	// 	log.Printf("not able to parse error returned %v", err)

	// 	return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	// }
	// return ctx.JSON(http.StatusOK, author)

}

// Delete an author by its id
func (h *Handler) DeleteAuthor(ctx echo.Context) error {
	//id := ctx.Param("id")

	return ctx.NoContent(http.StatusNotImplemented)
	// err := h.gateway.DeleteAuthor(ctx.Request().Context(), id)

	// if err != nil {
	// 	if e, ok := status.FromError(err); ok {
	// 		switch e.Code() {
	// 		case codes.NotFound:
	// 			return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
	// 		default:
	// 			log.Printf("DeleteAuthor failed: Err: %v\n", err)
	// 			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	// 		}
	// 	}

	// 	log.Printf("not able to parse error returned %v", err)

	// 	return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	// }

	//return ctx.NoContent(http.StatusOK)
}
