package models

import (
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty"`
}

func (d *UserDocument) ToModel() *models.User {
	return &models.User{
		ID:        d.ID.Hex(),
		Username:  d.Username,
		Email:     d.Email,
		FirstName: d.FirstName,
		LastName:  d.LastName,
	}
}
