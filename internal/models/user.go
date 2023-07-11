package models

import "errors"

var (
	ErrNotValidUser = errors.New("not valid user values")
)

type User struct {
	ID       int    `json:"id" db:"id" `
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" db:"phone"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}

type Auth struct {
	Username string `json:"username"`
}
