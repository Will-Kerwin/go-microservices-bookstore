package db

import (
	"context"

	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
)

type AuthorRepository interface {
	Add(ctx context.Context, author *models.Author) error
	Get(ctx context.Context) ([]*models.Author, error)
	GetById(ctx context.Context, id string) (*models.Author, error)
	Delete(ctx context.Context, id string) error
}

type BookRepository interface{}
