package service

import "net/http"

type ErrorCode int

const (
	ErrBadRequest ErrorCode = iota
	ErrForbidden
	ErrNotFound
	ErrUnauthorized
	ErrInternalServer
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func ErrorCodeToHTTPStatus(code ErrorCode) int {
	switch code {
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInternalServer:
		return http.StatusInternalServerError
	case ErrForbidden:
		return http.StatusForbidden
	case ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
