package models

import (
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/user"
)

var EmailRegex *regexp.Regexp = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

func BuildJwt(username string, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := &JwtCustomClaims{
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))

}

type JwtCustomClaims struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Roles    []user.UserRole `json:"roles"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	Token string `json:"token"`
}
