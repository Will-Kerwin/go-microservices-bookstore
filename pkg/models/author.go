package models

import (
	"time"

	"github.com/will-kerwin/go-microservice-bookstore/gen"
)

type Author struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

func AuthorToProto(a *Author) *gen.Author {
	return &gen.Author{
		Id:          a.ID,
		Name:        a.Name,
		DateOfBirth: a.DateOfBirth.Unix(),
	}
}

func AuthorsToProtos(a []*Author) []*gen.Author {

	var authors []*gen.Author = []*gen.Author{}

	for i := range a {
		authors = append(authors, AuthorToProto(a[i]))
	}

	return authors
}

func ProtoToAuthor(a *gen.Author) *Author {

	return &Author{
		ID:          a.Id,
		Name:        a.Name,
		DateOfBirth: time.Unix(a.DateOfBirth, 0),
	}
}

func ProtosToAuthors(a []*gen.Author) []*Author {
	var authors []*Author = []*Author{}

	for i := range a {
		authors = append(authors, ProtoToAuthor(a[i]))
	}

	return authors
}
