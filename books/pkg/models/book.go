package models

import (
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookDocument struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	AuthorId primitive.ObjectID `json:"authorId,omitempty" bson:"authorId,omitempty"`
	Synopsis string             `json:"synopsis"`
	ImageUrl string             `json:"imageUrl"`
	Genre    string             `json:"genre"`
}

func (d *BookDocument) ToModel() *models.Book {
	return &models.Book{
		ID:       d.ID.Hex(),
		Title:    d.Title,
		AuthorId: d.AuthorId.Hex(),
		Synopsis: d.Synopsis,
		ImageUrl: d.ImageUrl,
		Genre:    d.Genre,
	}
}
