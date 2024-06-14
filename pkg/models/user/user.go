package user

import (
	"github.com/will-kerwin/go-microservice-bookstore/gen"
)

type User struct {
	ID        string     `json:"_id,omitempty"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Roles     []UserRole `json:"roles"`
}

type UserRole string

const (
	Admin UserRole = "admin"
)

func UserToProto(user *User) *gen.User {
	return &gen.User{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ProtoToUser(user *gen.User) *User {
	return &User{
		ID:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
