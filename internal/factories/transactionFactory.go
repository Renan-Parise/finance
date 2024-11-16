package factories

import (
	"time"

	"github.com/Renan-Parise/finances/internal/entities"
)

func NewTransaction(userID int64, description string, category int, amount float64) *entities.Transaction {
	now := time.Now()
	return &entities.Transaction{
		UserID:      userID,
		CreatedAt:   now,
		UpdatedAt:   now,
		Description: description,
		Category:    category,
		Amount:      amount,
	}
}
