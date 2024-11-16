package container

import (
	"github.com/Renan-Parise/finances/internal/db"
	"github.com/Renan-Parise/finances/internal/repositories"
	"github.com/Renan-Parise/finances/internal/usecases"
)

type Container struct {
	TransactionRepository repositories.Transactionrepositories
	CategoryRepository    repositories.CategoryRepository

	TransactionUseCase usecases.TransactionUseCase
	CategoryUseCase    usecases.CategoryUseCase
}

func NewContainer() *Container {
	database := db.GetDB()

	transactionRepo := repositories.NewTransactionrepositories(database)
	categoryRepo := repositories.NewCategoryRepository(database)

	transactionUseCase := usecases.NewTransactionUseCase(transactionRepo)
	categoryUseCase := usecases.NewCategoryUseCase(categoryRepo)

	return &Container{
		TransactionRepository: transactionRepo,
		CategoryRepository:    categoryRepo,
		TransactionUseCase:    transactionUseCase,
		CategoryUseCase:       categoryUseCase,
	}
}
