package model

import "time"

type User struct {
	Email        string
	Username     string
	FirstName    string
	LastName     string
	Password     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
