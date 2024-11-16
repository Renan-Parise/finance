package usecases

import (
	"time"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/Renan-Parise/finances/internal/factories"
	"github.com/Renan-Parise/finances/internal/repositories"
)

type TransactionUseCase interface {
	CreateTransaction(userID int64, description string, category int, amount float64) error
	GetTransactions(userID int64) ([]*entities.Transaction, error)
	UpdateTransaction(transaction *entities.Transaction) error
	DeleteTransaction(userID int64, id int64) error
}

type transactionUseCase struct {
	transactionRepo repositories.Transactionrepositories
}

func NewTransactionUseCase(tr repositories.Transactionrepositories) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: tr,
	}
}

func (uc *transactionUseCase) CreateTransaction(userID int64, description string, category int, amount float64) error {
	transaction := factories.NewTransaction(userID, description, category, amount)
	return uc.transactionRepo.Create(transaction)
}

func (uc *transactionUseCase) GetTransactions(userID int64) ([]*entities.Transaction, error) {
	return uc.transactionRepo.GetAll(userID)
}

func (uc *transactionUseCase) UpdateTransaction(transaction *entities.Transaction) error {
	transaction.UpdatedAt = time.Now()
	return uc.transactionRepo.Update(transaction)
}

func (uc *transactionUseCase) DeleteTransaction(userID int64, id int64) error {
	return uc.transactionRepo.Delete(userID, id)
}
