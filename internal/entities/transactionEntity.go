package entities

import "time"

type Transaction struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Description string    `json:"description"`
	Category    int       `json:"category"`
	Amount      float64   `json:"amount"`
}
