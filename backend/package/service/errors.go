package service

import "fmt"

var (
	ErrNotFound        = fmt.Errorf("data not found")
	ErrAlreadyExists   = fmt.Errorf("data already exists")
	ErrInvalidToken    = fmt.Errorf("invalid token")
	ErrInvalidPassword = fmt.Errorf("invalid password")
	ErrInvalidInput    = fmt.Errorf("invalid input")
	ErrInputTooLong    = fmt.Errorf("input too long")
)
