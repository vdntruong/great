package apperror

import "errors"

var ( // TODO: improve there errors to have more context
	ErrUsernameExisted = errors.New("username existed")
	ErrEmailExisted    = errors.New("email existed")
)
