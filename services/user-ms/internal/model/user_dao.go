package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"gcommons/otel/trace"

	"user-ms/internal/pkg/apperror"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// DAOUser is an abstract interface to User of database mechanism, encapsulating all access to the data source.
type DAOUser struct {
	tracer trace.Tracer
	db     *sql.DB
}

func NewDAOUser(db *sql.DB, tracer trace.Tracer) *DAOUser {
	return &DAOUser{db: db, tracer: tracer}
}

func (u *DAOUser) Insert(c context.Context, email, username string, passwordHash string) (*User, error) {
	ctx, span := u.tracer.Start(c, "insert new user")
	defer span.End()

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
		span.RecordError(err)

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
		dbUser: dbUser{
			ID:           id,
			Email:        email,
			Username:     username,
			PasswordHash: passwordHash,
			CreatedAt:    createdAt,
			UpdatedAt:    createdAt,
		},
	}, nil
}

func (u *DAOUser) GetByEmail(c context.Context, email string) (*User, error) {
	ctx, span := u.tracer.Start(c, "get by email")
	defer span.End()

	var (
		query = `SELECT id, email, username, created_at, updated_at FROM users WHERE email = $1`
		row   = u.db.QueryRowContext(ctx, query, email)
	)

	var user User
	var err = row.Scan(&user.ID, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		span.RecordError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to scan user: %w", ErrUserNotFound)
	}

	return &user, nil
}

func (u *DAOUser) FindByEmail(c context.Context, email string) (bool, error) {
	ctx, span := u.tracer.Start(c, "find by email")
	defer span.End()

	var (
		query = `SELECT COUNT (email) FROM users WHERE email = $1`
		row   = u.db.QueryRowContext(ctx, query, email)

		total   int64
		err     = row.Scan(&total)
		founded = err == nil && total == 1
	)
	if errors.Is(err, sql.ErrNoRows) {
		return founded, nil
	}
	if err != nil {
		span.RecordError(err)
		return founded, fmt.Errorf("failed to scan total: %w", err)
	}

	return founded, nil
}

func (u *DAOUser) FindByUsername(c context.Context, username string) (bool, error) {
	ctx, span := u.tracer.Start(c, "find by username")
	defer span.End()

	var (
		query = `SELECT COUNT (username) FROM users WHERE username = $1`
		row   = u.db.QueryRowContext(ctx, query, username)

		total   int64
		err     = row.Scan(&total)
		founded = err == nil && total == 1
	)
	if errors.Is(err, sql.ErrNoRows) {
		return founded, nil
	}
	if err != nil {
		span.RecordError(err)
		return founded, fmt.Errorf("failed to scan total: %w", err)
	}

	return founded, nil
}
