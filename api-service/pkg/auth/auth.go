package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
)

var JwtConfig echojwt.Config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(models.JwtCustomClaims)
	},
	SigningKey: []byte("secret"),
	ErrorHandler: func(c echo.Context, err error) error {
		return c.JSON(http.StatusUnauthorized, models.ApiErrorResponse{"error": "user is not authenticated"})
	},
}
