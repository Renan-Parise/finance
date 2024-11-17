package usecases

import (
	"fmt"
	"io"
	"time"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/Renan-Parise/finances/internal/factories"
	"github.com/Renan-Parise/finances/internal/repositories"
	"github.com/Renan-Parise/finances/internal/utils/exporter"
)

type TransactionUseCase interface {
	FilterTransactions(userID int64, filter *entities.Filter) ([]*entities.Transaction, error)
	CreateTransaction(userID int64, description string, category int, amount float64) error
	ExportTransactions(userID int64, filter *entities.Filter, w io.Writer) error
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

func (uc *transactionUseCase) FilterTransactions(userID int64, filter *entities.Filter) ([]*entities.Transaction, error) {
	return uc.transactionRepo.Filter(userID, filter)
}

func (uc *transactionUseCase) ExportTransactions(userID int64, filter *entities.Filter, w io.Writer) error {
	transactions, err := uc.FilterTransactions(userID, filter)
	if err != nil {
		return err
	}

	exportContext := exporter.ExportContext{}

	switch filter.File {
	case "XLSX":
		exportContext.SetStrategy(&exporter.XLSXExporter{})
	case "PDF":
		exportContext.SetStrategy(&exporter.PDFExporter{})
	default:
		return fmt.Errorf("unsupported file type: %s", filter.File)
	}

	return exportContext.Execute(transactions, w)
}
