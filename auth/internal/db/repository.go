package db

import (
	"context"

	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
)

type AuthRepository interface {
	Add(ctx context.Context, user *models.CreateUserEvent) (*models.User, error)
	GetById(ctx context.Context, id string) (*models.User, error)
	Authenticate(ctx context.Context, username string, password string) (*models.User, error)
	ValidateUsernameUnique(ctx context.Context, username string) (bool, error)
}
