package models

import "time"

type CreateAuthorEvent struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type DeleteAuthorEvent struct {
	ID string `json:"_id"`
}
