package models

import "errors"

var ErrUnauthenticated = errors.New("not authenticated")
var ErrUserNotFound = errors.New("user not found")
