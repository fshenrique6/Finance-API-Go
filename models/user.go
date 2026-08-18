package models

import "time"

type User struct {
	ID		   int       `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Password   string    `json:"password,omitempty" binding:"required,min=6"`
	PasswordHash string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
}