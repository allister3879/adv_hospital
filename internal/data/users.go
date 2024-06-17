package data

import (
	"database/sql"

	db "d.hospital/pkg"
)

type User struct {
	email    string
	password string
}

type UserModel struct {
	db *sql.DB
}

func NewUserRepository() *UserModel {
	return &UserModel{
		db: db.DB,
	}
}

func (r *UserModel) RegisterUser(email, password string) (string, error) {
	err := r.db.QueryRow("INSERT INTO users (email, pass) VALUES ($1, $2) RETURNING email", email, password).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (r *UserModel) LoginUser(email, password string) (string, error) {
	err := r.db.QueryRow("GET FROM users * where email $1", email, password).Scan(&email, &password)
	if err != nil {
		return "", err
	}
	return "logined", nil
}
