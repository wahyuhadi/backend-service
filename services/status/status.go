package error

import "net/http"

type ErrorWrapper struct {
	Code int
	err  error
}

func (e *ErrorWrapper) Error() string {
	return e.err.Error()
}

func (e *ErrorWrapper) Wrap(err error) error {
	e.err = err
	return e
}

var (
	StatusOK                  = ErrorWrapper{Code: http.StatusOK}
	StatusCreated             = ErrorWrapper{Code: http.StatusCreated}
	StatusAccepted            = ErrorWrapper{Code: http.StatusAccepted}
	StatusNoContent           = ErrorWrapper{Code: http.StatusNoContent}
	StatusBadRequest          = ErrorWrapper{Code: http.StatusBadRequest}
	StatusUnauthorized        = ErrorWrapper{Code: http.StatusUnauthorized}
	StatusForbidden           = ErrorWrapper{Code: http.StatusForbidden}
	StatusNotFound            = ErrorWrapper{Code: http.StatusNotFound}
	StatusConflict            = ErrorWrapper{Code: http.StatusConflict}
	StatusInternalServerError = ErrorWrapper{Code: http.StatusInternalServerError}
	StatusNotImplemented      = ErrorWrapper{Code: http.StatusNotImplemented}
	StatusServiceUnavailable  = ErrorWrapper{Code: http.StatusServiceUnavailable}
	StatusTooManyRequests     = ErrorWrapper{Code: http.StatusTooManyRequests}
)
