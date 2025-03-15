package doiterr

import "errors"

var ErrEmptyToken = errors.New("empty token in config")
var ErrEmptyDbURI = errors.New("empty db uri in config")
var ErrEmptyParameters = errors.New("empty parameters")
var ErrEmptyEndpoint = errors.New("empty endpoint")

type Error struct {
	err   error
	cause error
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.cause
}

func WrapError(err error, cause error) error {
	return &Error{
		err:   err,
		cause: cause,
	}
}
