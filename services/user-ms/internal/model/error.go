package model

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")

	ErrUserNotFound = fmt.Errorf("user %w", ErrNotFound)

	ErrEmailRequired    = fmt.Errorf("user: email is required")
	ErrUsernameRequired = fmt.Errorf("user: username is required")
	ErrPasswordRequired = fmt.Errorf("user: password is required")
	ErrPasswordHash     = fmt.Errorf("user: could not gen password hash")
)
