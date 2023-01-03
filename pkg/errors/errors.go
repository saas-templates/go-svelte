package errors

import (
	"errors"
	"fmt"
	"strings"
)

// Common timer domain errors. Use `ErrX.WithCausef()` to clone and add context.
var (
	ErrInvalid     = Error{Code: "bad_request", Message: "Request is not valid"}
	ErrNotFound    = Error{Code: "not_found", Message: "Requested entity not found"}
	ErrConflict    = Error{Code: "conflict", Message: "An entity with conflicting identifier exists"}
	ErrForbidden   = Error{Code: "forbidden", Message: "You are not authorised for the requested action"}
	ErrInternal    = Error{Code: "internal_error", Message: "Some unexpected error occurred"}
	ErrUnsupported = Error{Code: "unsupported", Message: "Requested feature is not supported"}
)

// E converts any given error to the Error type. Unknown are converted
// to ErrInternal.
func E(err error) Error {
	if e, ok := err.(Error); ok {
		return e
	}
	return ErrInternal.WithCausef(err.Error())
}

// Error represents any error returned by the components along with any
// relevant context.
type Error struct {
	Code    string `json:"code"`
	Cause   string `json:"cause,omitempty"`
	Message string `json:"message"`
}

// WithCausef returns clone of err with the cause added. Use this when
// you need to provide description of the underlying technical root-cause
// which may be written in log for debugging purposes. Cause will be shown
// to the user only when the Message is empty.
func (err Error) WithCausef(format string, args ...interface{}) Error {
	cloned := err
	cloned.Cause = fmt.Sprintf(format, args...)
	return cloned
}

// WithMsgf returns a clone of the error with message set. Use this when
// you need to provide a custom message that should be shown to the user.
// If the message is set to empty string, cause will be displayed to the
// user.
func (err Error) WithMsgf(format string, args ...interface{}) Error {
	cloned := err
	cloned.Message = fmt.Sprintf(format, args...)
	return cloned
}

// Is checks if 'other' is of type Error and has the same code.
// See https://blog.golang.org/go1.13-errors.
func (err Error) Is(other error) bool {
	if oe, ok := other.(Error); ok {
		return oe.Code == err.Code
	}

	// unknown error types are considered as internal errors.
	return err.Code == ErrInternal.Code
}

func (err Error) Error() string {
	if err.Message != "" {
		return strings.ToLower(err.Message)
	}
	return fmt.Sprintf("%s: %s", err.Code, err.Cause)
}

// Errorf returns a formatted error similar to `fmt.Errorf` but uses the
// Error type defined in this package. returned value is equivalent to
// ErrInternal (i.e., errors.Is(retVal, ErrInternal) = true).
func Errorf(format string, args ...interface{}) error {
	return ErrInternal.WithMsgf(format, args...)
}

// Is returns true if 'err' is equivalent to the 'target' error.
// This function is a convenience shortcut for the errors.Is().
func Is(err, target error) bool { return errors.Is(err, target) }

// New returns a new error equivalent to ErrInternal.
// This function is a convenience shortcut for the errors.New().
func New(msg string) error { return errors.New(msg) }
