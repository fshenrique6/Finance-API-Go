package models

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"`     // "income" ou "expense"
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}