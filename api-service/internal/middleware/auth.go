package middleware

import (
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/user"
)

func UseAdminOrSameUserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqId := c.Param("id")
		jwt := c.Get("user").(*jwt.Token)
		claims := jwt.Claims.(*models.JwtCustomClaims)
		if claims.ID != reqId || slices.Contains(claims.Roles, user.Admin) {
			return c.JSON(http.StatusUnauthorized, models.ApiErrorResponse{"error": "user id does not match the requested id"})
		}

		return nil
	}
}
