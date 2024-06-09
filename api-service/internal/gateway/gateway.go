package gateway

import (
	"context"

	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
)

type AuthorGateway interface {
	Get(ctx context.Context) ([]*models.Author, error)
	GetById(ctx context.Context, id string) (*models.Author, error)
}
