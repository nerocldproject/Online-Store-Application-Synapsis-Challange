package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UserID    int        `json:"user_id"`
	UserName  string     `json:"user_name"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Claims struct {
	jwt.RegisteredClaims
	Username string
}