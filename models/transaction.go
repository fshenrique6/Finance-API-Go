package models

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	Description string    `json:"description" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=income expense"`
	Category    string    `json:"category" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}