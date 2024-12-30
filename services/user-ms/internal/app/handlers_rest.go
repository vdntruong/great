package app

import (
	"context"
	"errors"
	"net/http"

	"user-ms/internal/dto"
	"user-ms/internal/model"
	"user-ms/internal/pkg/apperror"

	"golang.org/x/sync/errgroup"
)

func (app *Application) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserReq
	if err := app.Decode(r, &req); err != nil {
		app.respondBadRequest(w, err)
		return
	}

	var ctx = r.Context()

	var eg errgroup.Group
	eg.Go(func() error {
		return app.shouldNotHaveUsername(ctx, req.Username)
	})
	eg.Go(func() error {
		return app.shouldNotHaveEmail(ctx, req.Email)
	})
	if err := eg.Wait(); err != nil {
		app.respondError(w, r, err)
		return
	}

	// handle validate state (valid states), and prepare for persistence
	baseUser, err := model.NewUser(req.Email, req.Username, req.Password)
	if err != nil {
		app.respondBadRequest(w, err)
		return
	}

	user, err := app.userRepo.Insert(ctx, baseUser.Email, baseUser.Username, baseUser.PasswordHash)
	if err != nil {
		app.respondError(w, r, err)
		return
	}

	resp := dto.UserRes{
		Email:    user.Email,
		Username: user.Username,
	}
	app.respondCreated(w, resp)
}

func (app *Application) shouldNotHaveUsername(ctx context.Context, username string) error {
	founded, err := app.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return err
	}
	if founded {
		return apperror.ErrUsernameExisted
	}
	return nil
}

func (app *Application) shouldNotHaveEmail(ctx context.Context, email string) error {
	founded, err := app.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if founded {
		return apperror.ErrEmailExisted
	}
	return nil
}

func (app *Application) HandleProfile(w http.ResponseWriter, r *http.Request) {
	var userID = r.Header.Get("X-User-ID")
	if userID == "" {
		app.respondBadRequest(w, apperror.ErrUserIDRequired)
		return
	}

	user, err := app.userRepo.GetByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			app.respondNotFound(w)
			return
		}
		app.respondError(w, r, err)
		return
	}

	resp := dto.ConvertToUserRes(*user)
	app.respondOK(w, resp)
}
