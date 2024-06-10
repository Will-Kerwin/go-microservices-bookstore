package gateway

import (
	"context"

	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
)

type AuthorGateway interface {
	Get(ctx context.Context) ([]*models.Author, error)
	GetById(ctx context.Context, id string) (*models.Author, error)
}

type BookGateway interface {
	Get(ctx context.Context, title string, authorId string, genre string) ([]*models.Book, error)
	GetById(ctx context.Context, id string) (*models.Book, error)
}
