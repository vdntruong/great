package models

import (
	"time"
)

type User struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password_hash"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u User) GetID() interface{} {
	return u.ID
}

func (u User) SetID(id interface{}) {
	u.ID = id.(string)
}

func (u User) Validate() error {
	return nil
}
