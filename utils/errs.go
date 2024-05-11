package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)


var ErrBadInput = errors.New("Bad input")
var ErrBadToken = errors.New("Bad token")
var ErrNoToken = errors.New("No token provided")
var ErrNotAuthorized = errors.New("Not authorized to access resource")
var ErrNotFound = errors.New("Resource not found")
var ErrInternalServer = errors.New("Internal server error")

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func GetHttpErrorFromError(err error, resourceName string) render.Renderer {
	if err == ErrBadInput {
		return constructErr(err, 400, "Invalid request.")
	} else if err == ErrNoToken {
		return constructErr(err, 401, "No token provided.")
	} else if err == ErrNotAuthorized {
		return constructErr(err, 403, fmt.Sprintf("Not authorized to access %s.", resourceName))
	} else if err == ErrNotFound {
		return constructErr(err, 404, fmt.Sprintf("%s not found.", resourceName))
	} else if err == ErrBadToken {
		return constructErr(err, 401, "Invalid token.")
	} else {
		return constructErr(err, 500, "Server Error")
	}
}

func constructErr(err error, status int, statusText string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: status,
		StatusText:     statusText,
		ErrorText:      err.Error(),
	}
}
