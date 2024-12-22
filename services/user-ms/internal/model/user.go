package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"user-ms/internal/pkg/apperror"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID           string `json:"id" db:"id" cql:"id"`
	Email        string `json:"email" db:"email" cql:"email"`
	Username     string `json:"username" db:"username" cql:"username"`
	PasswordHash string `json:"-" db:"password_hash" cql:"password_hash"`

	CreatedAt time.Time `json:"created_at" db:"created_at" cql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" cql:"updated_at"`
}

// UserModel is like a DAO
type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}
func (u *UserModel) Insert(ctx context.Context, email, username string, passwordHash string) (*User, error) {
	var (
		id        = uuid.NewString()
		createdAt = time.Now()

		query = `
            INSERT INTO users (id, email, username, password_hash, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $5)
        `

		_, err = u.db.ExecContext(ctx, query, id, email, username, passwordHash, createdAt)
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			// Check for unique violation
			if pqErr.Code == "23505" {
				if strings.Contains(pqErr.Constraint, "email") {
					return nil, fmt.Errorf("email already exists: %w, %w", apperror.ErrEmailExisted, err)
				}
				if strings.Contains(pqErr.Constraint, "username") {
					return nil, fmt.Errorf("username already exists: %w, %w", apperror.ErrUsernameExisted, err)
				}
			}
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &User{
		ID:           id,
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}, nil
}

func (u *UserModel) GetByEmail(ctx context.Context, email string) (*User, bool, error) {
	var (
		query = `SELECT id, email, username, created_at, updated_at FROM users WHERE email = $1`
		row   = u.db.QueryRowContext(ctx, query, email)
	)

	var user User
	var err = row.Scan(&user.ID, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("failed to scan user: %w", err)
	}

	return &user, true, nil
}
