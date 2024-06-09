package models

import (
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorDocument struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	DateOfBirth primitive.DateTime `json:"age"`
}

func (d *AuthorDocument) ToModel() *models.Author {
	return &models.Author{
		ID:          d.ID.Hex(),
		Name:        d.Name,
		DateOfBirth: d.DateOfBirth.Time(),
	}
}
