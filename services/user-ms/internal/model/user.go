package model

import (
	"fmt"
	"time"

	gpassword "gcommons/password"
)

type dbUser struct {
	ID           string `json:"id" db:"id" cql:"id"`
	Email        string `json:"email" db:"email" cql:"email"`
	Username     string `json:"username" db:"username" cql:"username"`
	PasswordHash string `json:"-" db:"password_hash" cql:"password_hash"`

	CreatedAt time.Time `json:"created_at" db:"created_at" cql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" cql:"updated_at"`
}

// User is the Value Object for dbUser, aka BaseUser
type User struct {
	dbUser
}

func NewUser(email, username, password string) (*User, error) {
	if err := validateUserFields(email, username, password); err != nil {
		return nil, err
	}

	passwordHash, err := gpassword.Hash(password)
	if err != nil {
		return nil, fmt.Errorf("%w (%w)", ErrPasswordHash, err)
	}

	return &User{
		dbUser: dbUser{
			Email:        email,
			Username:     username,
			PasswordHash: passwordHash,
		},
	}, nil
}

func validateUserFields(email, username, password string) error {
	if email == "" {
		return ErrEmailRequired
	}
	if username == "" {
		return ErrUsernameRequired
	}
	if password == "" {
		return ErrPasswordRequired
	}

	return nil
}
