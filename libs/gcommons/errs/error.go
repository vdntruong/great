package errs

import (
	"fmt"
	"net/http"
)

type Code int

type Error struct {
	err  error
	code Code

	statusCode int
	statusText string
}

func (e Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return e.GetStatusText()
}

func New(code Code) Error {
	return Error{
		code: code,
	}
}

func NewCustomError(statusCode int, statusText string) *Error {
	if !isHTTPStatusCode(statusCode) {
		statusCode = http.StatusInternalServerError
	}

	if statusText == "" {
		statusText = http.StatusText(statusCode)
	}

	return &Error{
		statusCode: statusCode,
		statusText: statusText,
	}
}

func isHTTPStatusCode(code int) bool {
	validRanges := [][]int{
		{http.StatusContinue, http.StatusEarlyHints},                               // 1xx Informational
		{http.StatusOK, http.StatusIMUsed},                                         // 2xx Success
		{http.StatusMultipleChoices, http.StatusPermanentRedirect},                 // 3xx Redirection
		{http.StatusBadRequest, http.StatusUnavailableForLegalReasons},             // 4xx Client Error
		{http.StatusInternalServerError, http.StatusNetworkAuthenticationRequired}, // 5xx Server Error
	}

	// Check if the code falls within any of the valid ranges
	for _, rangeValues := range validRanges {
		if code >= rangeValues[0] && code <= rangeValues[1] {
			return true
		}
	}

	return false
}

func (e Error) mapHTTPStatusCode() int {
	if e.code == 0 {
		return http.StatusInternalServerError
	}

	if statusCode, ok := errResponseHTTPStatusCodeMapper[e.code]; ok {
		return statusCode
	}
	return http.StatusInternalServerError
}

func (e Error) GetStatusCode() int {
	if e.statusCode != 0 {
		return e.statusCode
	}

	return e.mapHTTPStatusCode()
}

func (e Error) mapHTTPStatusText() string {
	if e.code == 0 {
		return http.StatusText(http.StatusInternalServerError)
	}

	if statusText, ok := errResponseHTTPStatusTextMapper[e.code]; ok {
		return statusText
	}

	return fmt.Sprintf("%s (%d)", http.StatusText(http.StatusInternalServerError), e.code)
}

func (e Error) GetStatusText() string {
	if e.statusText != "" {
		return e.statusText
	}

	return e.mapHTTPStatusText()
}

var (
	errResponseHTTPStatusTextMapper = map[Code]string{
		http.StatusInternalServerError: http.StatusText(http.StatusInternalServerError),
	}

	errResponseHTTPStatusCodeMapper = map[Code]int{
		http.StatusInternalServerError: http.StatusInternalServerError,
	}
)

func OverrideStatusTextsMapper(m map[Code]string) {
	errResponseHTTPStatusTextMapper = m
}

func OverrideStatusCodesMapper(m map[Code]int) {
	errResponseHTTPStatusCodeMapper = m
}
