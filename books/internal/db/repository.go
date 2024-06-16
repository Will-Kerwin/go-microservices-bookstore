package db

import (
	"context"

	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/events"
)

type AuthorRepository interface {
	Add(ctx context.Context, author *models.Author) (*models.Author, error)
	Get(ctx context.Context) ([]*models.Author, error)
	GetById(ctx context.Context, id string) (*models.Author, error)
	Delete(ctx context.Context, id string) error
}

type BookRepository interface {
	Add(ctx context.Context, book *models.Book) (*models.Book, error)
	Get(ctx context.Context, title string, authorId string, genre string) ([]*models.Book, error)
	GetById(ctx context.Context, id string) (*models.Book, error)
	Update(ctx context.Context, id string, updatedBook *events.UpdateBookEventData) error
	Delete(ctx context.Context, id string) error
}
