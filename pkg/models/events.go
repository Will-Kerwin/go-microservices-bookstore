package models

import "time"

type CreateAuthorEvent struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type DeleteAuthorEvent struct {
	ID string `json:"_id"`
}

type CreateBookEvent struct {
	Title    string `json:"title"`
	AuthorId string `json:"authorId"`
	Synopsis string `json:"synopsis"`
	ImageUrl string `json:"imageUrl"`
	Genre    string `json:"genre"`
}

type DeleteBookEvent struct {
	ID string `json:"_id"`
}

type UpdateBookEvent struct {
	ID   string              `json:"_id"`
	Data UpdateBookEventData `json:"data"`
}

type UpdateBookEventData struct {
	Title    string `json:"title,omitempty"`
	AuthorId string `json:"authorId,omitempty"`
	Synopsis string `json:"synopsis,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Genre    string `json:"genre,omitempty"`
}
