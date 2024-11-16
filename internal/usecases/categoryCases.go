package usecases

import (
	"time"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/Renan-Parise/finances/internal/repositories"
)

type CategoryUseCase interface {
	CreateCategory(userID int64, name string) error
	GetCategories(userID int64) ([]*entities.Category, error)
	DeleteCategory(userID int64, id int) error
}

type categoryUseCase struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryUseCase(cr repositories.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{categoryRepo: cr}
}

func (uc *categoryUseCase) CreateCategory(userID int64, name string) error {
	category := &entities.Category{
		UserID:    userID,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return uc.categoryRepo.Create(category)
}

func (uc *categoryUseCase) GetCategories(userID int64) ([]*entities.Category, error) {
	return uc.categoryRepo.GetAll(userID)
}

func (uc *categoryUseCase) DeleteCategory(userID int64, id int) error {
	return uc.categoryRepo.Delete(userID, id)
}
