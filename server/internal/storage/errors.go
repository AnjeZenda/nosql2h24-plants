package storage

import "errors"

// User errors
var (
	ErrUserInvalidId = errors.New("invalid id")
	ErrUserNotFound  = errors.New("user not found")
)
