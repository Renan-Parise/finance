package entities

import "time"

type Transaction struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	DateAdded   time.Time `json:"date_added"`
	DateEdited  time.Time `json:"date_edited"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
}
