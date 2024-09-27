package storage

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrInsufficentPermissions = errors.New("insufficent permsissions")
)
