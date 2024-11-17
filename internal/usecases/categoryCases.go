package usecases

import (
	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/Renan-Parise/finances/internal/errors"
	"github.com/Renan-Parise/finances/internal/factories"
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
	exists, err := uc.categoryRepo.ExistsByName(userID, name)
	if err != nil {
		return errors.NewServiceError("error checking if category name exists: " + err.Error())
	}

	if exists {
		return errors.NewValidationError(name, "the given category name already exists: "+name)
	}

	category := factories.NewCategory(userID, name)
	return uc.categoryRepo.Create(category)
}

func (uc *categoryUseCase) GetCategories(userID int64) ([]*entities.Category, error) {
	return uc.categoryRepo.GetAll(userID)
}

func (uc *categoryUseCase) DeleteCategory(userID int64, id int) error {
	return uc.categoryRepo.Delete(userID, id)
}
