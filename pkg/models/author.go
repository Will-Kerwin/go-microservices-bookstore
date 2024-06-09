package models

import (
	"time"
)

type Author struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"age"`
}
