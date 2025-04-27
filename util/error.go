package util

import "errors"

var (
	ErrInvalidRequest      = errors.New("Invalid request")
	ErrInternalServerError = errors.New("Internal server error")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
	ErrInsufficientBalance = errors.New("Insufficient balance")
	ErrNasabahNotFound     = errors.New("Nasabah not found")
	ErrNasabahAlreadyExist = errors.New("Nasabah already exist")
)
