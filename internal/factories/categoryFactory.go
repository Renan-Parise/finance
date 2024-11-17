package factories

import (
	"time"

	"github.com/Renan-Parise/finances/internal/entities"
)

func NewCategory(userID int64, name string) *entities.Category {
	now := time.Now()
	return &entities.Category{
		UserID:    userID,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
