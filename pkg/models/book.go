package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	AuthorID primitive.ObjectID `json:"authorId"`
	Synopsis string             `json:"synopsis"`
	ImageUrl string             `json:"imageUrl"`
	Genre    string             `json:"genre"`
}
