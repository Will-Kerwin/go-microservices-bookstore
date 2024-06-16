package db

import (
	"context"

	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/events"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/user"
)

type AuthRepository interface {
	Add(ctx context.Context, user *events.CreateUserEvent) (*user.User, error)
	GetById(ctx context.Context, id string) (*user.User, error)
	Authenticate(ctx context.Context, username string, password string) (*user.User, error)
	ValidateUsernameUnique(ctx context.Context, username string) (bool, error)
	Update(ctx context.Context, id string, updateData *events.UpdateUserEventData) error
}
