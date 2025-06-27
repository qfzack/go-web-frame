package model

import "fmt"

type BusinessError struct {
	Code    string
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

var (
	ErrUserNotFound = &BusinessError{Code: "USER_NOT_FOUND", Message: "User not found"}
	ErrUserExists   = &BusinessError{Code: "USER_EXISTS", Message: "User already exists"}
)
