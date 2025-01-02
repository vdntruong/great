package models

import (
	"time"
)

type User struct {
	ID       string `db:"id"            json:"id"`
	Email    string `db:"email"         json:"email"`
	Username string `db:"username"      json:"username"`
	Password string `db:"password_hash" json:"-"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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
