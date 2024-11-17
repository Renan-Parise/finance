package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/Renan-Parise/finances/internal/errors"
)

type Transactionrepositories interface {
	Create(transaction *entities.Transaction) error
	GetAll(userID int64) ([]*entities.Transaction, error)
	GetByID(userID int64, id int64) (*entities.Transaction, error)
	Update(transaction *entities.Transaction) error
	Delete(userID int64, id int64) error
	Filter(userID int64, filter *entities.Filter) ([]*entities.Transaction, error)
}

type transactionrepositories struct {
	db *sql.DB
}

func NewTransactionrepositories(db *sql.DB) Transactionrepositories {
	return &transactionrepositories{
		db: db,
	}
}

func (r *transactionrepositories) Create(transaction *entities.Transaction) error {
	query := `INSERT INTO transactions (userId, createdAt, updatedAt, description, category, amount)
              VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errors.NewQueryError("error preparing query: " + err.Error())
	}
	defer stmt.Close()

	res, err := stmt.Exec(transaction.UserID, transaction.CreatedAt, transaction.UpdatedAt,
		transaction.Description, transaction.Category, transaction.Amount)
	if err != nil {
		return errors.NewQueryError("error executing query: " + err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		return errors.NewQueryError("error getting last insert ID: " + err.Error())
	}
	transaction.ID = id
	return nil
}

func (r *transactionrepositories) GetAll(userID int64) ([]*entities.Transaction, error) {
	query := `SELECT id, userId, createdAt, updatedAt, description, category, amount 
              FROM transactions 
              WHERE userId = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.NewQueryError("error executing query: " + err.Error())
	}
	defer rows.Close()

	var transactions []*entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.CreatedAt,
			&transaction.UpdatedAt, &transaction.Description, &transaction.Category, &transaction.Amount)
		if err != nil {
			return nil, errors.NewQueryError("error scanning row: " + err.Error())
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}

func (r *transactionrepositories) GetByID(userID int64, id int64) (*entities.Transaction, error) {
	query := `SELECT id, userId, createdAt, updatedAt, description, category, amount 
              FROM transactions 
              WHERE id = ? AND userId = ?`
	row := r.db.QueryRow(query, id, userID)
	var transaction entities.Transaction
	err := row.Scan(&transaction.ID, &transaction.UserID, &transaction.CreatedAt,
		&transaction.UpdatedAt, &transaction.Description, &transaction.Category, &transaction.Amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.NewQueryError("error scanning row: " + err.Error())
	}
	return &transaction, nil
}

func (r *transactionrepositories) Update(transaction *entities.Transaction) error {
	query := `UPDATE transactions SET updatedAt = ?, description = ?, category = ?, amount = ? 
              WHERE id = ? AND userId = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errors.NewQueryError("error preparing query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(transaction.UpdatedAt, transaction.Description, transaction.Category,
		transaction.Amount, transaction.ID, transaction.UserID)
	return err
}

func (r *transactionrepositories) Delete(userID int64, id int64) error {
	query := `DELETE FROM transactions WHERE id = ? AND userId = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errors.NewQueryError("error preparing query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, userID)
	return err
}

func (r *transactionrepositories) Filter(userID int64, filter *entities.Filter) ([]*entities.Transaction, error) {
	query := `
		SELECT id, userId, createdAt, updatedAt, description, category, amount
		FROM transactions
		WHERE userId = ?
	`
	args := []interface{}{userID}

	if filter.Category != 0 {
		query += " AND category = ?"
		args = append(args, filter.Category)
	}

	if filter.Search != "" {
		query += " AND LOWER(description) LIKE LOWER(?)"
		args = append(args, "%"+filter.Search+"%")
	}

	if filter.From != "" {
		query += " AND createdAt >= ?"
		args = append(args, filter.From)
	}

	if filter.To != "" {
		query += " AND createdAt <= ?"
		args = append(args, filter.To)
	}

	if filter.Field != "" {
		allowedFields := map[string]bool{
			"createdAt":   true,
			"description": true,
			"amount":      true,
		}
		if !allowedFields[filter.Field] {
			return nil, fmt.Errorf("invalid field for sorting: %s", filter.Field)
		}

		order := "ASC"
		if strings.ToUpper(filter.Order) == "DESC" {
			order = "DESC"
		}

		query += fmt.Sprintf(" ORDER BY %s %s", filter.Field, order)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, errors.NewQueryError("error executing query: " + err.Error())
	}
	defer rows.Close()

	var transactions []*entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.CreatedAt,
			&transaction.UpdatedAt, &transaction.Description, &transaction.Category, &transaction.Amount)
		if err != nil {
			return nil, errors.NewQueryError("error scanning row: " + err.Error())
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}
