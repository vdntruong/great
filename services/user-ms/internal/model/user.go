package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

type User struct {
	ID       string `json:"id" db:"id" cql:"id"`
	Email    string `json:"email" db:"email" cql:"email"`
	Username string `json:"username" db:"username" cql:"username"`
	Password string `json:"-" db:"password_hash" cql:"password_hash"`

	CreatedAt time.Time `json:"created_at" db:"created_at" cql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" cql:"updated_at"`
}

// UserModel is like a DAO
type UserModel struct {
	Session *gocql.Session
}

func (u *UserModel) InsertUser(ctx context.Context, email string, password string) (*User, error) {
	var (
		id        = gocql.TimeUUID().String()
		createdAt = time.Now()

		query = `INSERT INTO users (id, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
		args  = []interface{}{id, email, password, createdAt, createdAt}

		err = u.Session.Query(query, args...).WithContext(ctx).Exec()
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}

func (u *UserModel) GetUserByEmail(ctx context.Context, email string) (*User, bool, error) {
	var (
		query = `SELECT id, email, name, created_at, updated_at FROM users WHERE email = ?`
		args  = []interface{}{email}

		user User
		err  = u.Session.Query(query, args...).WithContext(ctx).
			Scan(&user.ID, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	)

	if errors.Is(err, gocql.ErrNotFound) {
		return nil, false, nil
	}

	return &user, true, err
}
