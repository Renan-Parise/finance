package repositories

import (
	"database/sql"
	"errors"

	"github.com/Renan-Parise/finances/internal/entities"
)

type Transactionrepositories interface {
	Create(transaction *entities.Transaction) error
	GetAll(userID int64) ([]*entities.Transaction, error)
	GetByID(userID int64, id int64) (*entities.Transaction, error)
	Update(transaction *entities.Transaction) error
	Delete(userID int64, id int64) error
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
	query := `INSERT INTO transactions (user_id, date_added, date_edited, description, category, amount)
              VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(transaction.UserID, transaction.DateAdded, transaction.DateEdited,
		transaction.Description, transaction.Category, transaction.Amount)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	transaction.ID = id
	return nil
}

func (r *transactionrepositories) GetAll(userID int64) ([]*entities.Transaction, error) {
	query := `SELECT id, user_id, date_added, date_edited, description, category, amount 
              FROM transactions 
              WHERE user_id = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.DateAdded,
			&transaction.DateEdited, &transaction.Description, &transaction.Category, &transaction.Amount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}

func (r *transactionrepositories) GetByID(userID int64, id int64) (*entities.Transaction, error) {
	query := `SELECT id, user_id, date_added, date_edited, description, category, amount 
              FROM transactions 
              WHERE id = ? AND user_id = ?`
	row := r.db.QueryRow(query, id, userID)
	var transaction entities.Transaction
	err := row.Scan(&transaction.ID, &transaction.UserID, &transaction.DateAdded,
		&transaction.DateEdited, &transaction.Description, &transaction.Category, &transaction.Amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionrepositories) Update(transaction *entities.Transaction) error {
	query := `UPDATE transactions SET date_edited = ?, description = ?, category = ?, amount = ? 
              WHERE id = ? AND user_id = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(transaction.DateEdited, transaction.Description, transaction.Category,
		transaction.Amount, transaction.ID, transaction.UserID)
	return err
}

func (r *transactionrepositories) Delete(userID int64, id int64) error {
	query := `DELETE FROM transactions WHERE id = ? AND user_id = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, userID)
	return err
}
