package apperror

import (
	"net/http"
	"sync"

	"gcommons/errs"
)

var (
	ErrUsernameExisted = errs.New(codeEmailExisted)
	ErrEmailExisted    = errs.New(codeUsernameExisted)
)

var (
	codeEmailExisted    errs.Code = 10_001
	codeUsernameExisted errs.Code = 10_002
)

var (
	errStatusCodeMapper = map[errs.Code]int{
		codeEmailExisted:    http.StatusConflict,
		codeUsernameExisted: http.StatusConflict,
	}

	errStatusTextMapper = map[errs.Code]string{
		codeEmailExisted:    "email already existed",
		codeUsernameExisted: "username already existed",
	}
)

var once sync.Once

func init() {
	once.Do(func() {
		errs.OverrideStatusCodesMapper(errStatusCodeMapper)
		errs.OverrideStatusTextsMapper(errStatusTextMapper)
	})
}
