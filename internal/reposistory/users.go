package repository

import (
	"GoProject/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
	GetUserId(id int) (models.User, error)
}

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) CreateUser(user models.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, phone, password_hash) VALUES (:name, :username, :phone, :password_hash)", "users")
	result, err := u.db.NamedExec(query, user)
	if err != nil {
		return 0, fmt.Errorf("repo: create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("repo: create user: %w", err)
	}

	return int(id), nil
}

func (u *User) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username= ?s AND password_hash=?", "users")
	err := u.db.Get(&user, query, username, password)
	if err != nil {
		return models.User{}, fmt.Errorf("repo: get user: %w", err)
	}
	return user, nil
}

func (u *User) GetUserId(id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", "users")
	err := u.db.Get(&user, query, id)
	if err != nil {
		return models.User{}, fmt.Errorf("repo: get user id: %w", err)
	}
	return user, err
}
