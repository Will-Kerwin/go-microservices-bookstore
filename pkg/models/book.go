package models

import (
	"github.com/will-kerwin/go-microservice-bookstore/gen"
)

type Book struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string `json:"title"`
	AuthorId string `json:"authorId"`
	Synopsis string `json:"synopsis"`
	ImageUrl string `json:"imageUrl"`
	Genre    string `json:"genre"`
}

func BookToProto(b *Book) *gen.Book {
	return &gen.Book{
		Id:       b.ID,
		Title:    b.Title,
		AuthorId: b.AuthorId,
		Synopsis: b.Synopsis,
		ImageUrl: b.ImageUrl,
		Genre:    b.Genre,
	}
}

func BooksToProtos(b []*Book) []*gen.Book {

	var Books []*gen.Book = []*gen.Book{}

	for i := range b {
		Books = append(Books, BookToProto(b[i]))
	}

	return Books
}

func ProtoToBook(b *gen.Book) *Book {

	return &Book{
		ID:       b.Id,
		Title:    b.Title,
		AuthorId: b.AuthorId,
		Synopsis: b.Synopsis,
		ImageUrl: b.ImageUrl,
		Genre:    b.Genre,
	}
}

func ProtosToBooks(b []*gen.Book) []*Book {
	var Books []*Book = []*Book{}

	for i := range b {
		Books = append(Books, ProtoToBook(b[i]))
	}

	return Books
}
