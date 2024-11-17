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

	StatisticsRepository repositories.StatisticsRepository
	StatisticsUseCase    usecases.StatisticsUseCase
}

func NewContainer() *Container {
	database := db.GetDB()

	transactionRepo := repositories.NewTransactionrepositories(database)
	statisticsRepo := repositories.NewStatisticsRepository(database)
	categoryRepo := repositories.NewCategoryRepository(database)

	transactionUseCase := usecases.NewTransactionUseCase(transactionRepo)
	statisticsUseCase := usecases.NewStatisticsUseCase(statisticsRepo)
	categoryUseCase := usecases.NewCategoryUseCase(categoryRepo)

	return &Container{
		TransactionRepository: transactionRepo,
		CategoryRepository:    categoryRepo,
		StatisticsRepository:  statisticsRepo,

		StatisticsUseCase:  statisticsUseCase,
		TransactionUseCase: transactionUseCase,
		CategoryUseCase:    categoryUseCase,
	}
}
