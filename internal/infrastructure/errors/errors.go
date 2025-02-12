package errors

import "errors"

var (
	// ErrInvalidInput    = errors.New("invalid input")
	// ErrNotFound        = errors.New("item not found")
	// ErrInternal        = errors.New("internal server error")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInternal          = errors.New("internal server error")
	ErrInsufficientFound = errors.New("insufficient funds")
	ErrMerchNotFound     = errors.New("merch item not found")
)
