package repositories

import (
	"database/sql"

	"github.com/Renan-Parise/finances/internal/entities"
)

type CategoryRepository interface {
	Create(category *entities.Category) error
	GetAll(userID int64) ([]*entities.Category, error)
	Delete(userID int64, id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *entities.Category) error {
	query := `INSERT INTO categories (userId, name, createdAt, updatedAt) VALUES (?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(category.UserID, category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = int(id)
	return nil
}

func (r *categoryRepository) GetAll(userID int64) ([]*entities.Category, error) {
	query := `SELECT id, userId, name, createdAt, updatedAt FROM categories WHERE userId = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entities.Category
	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.ID, &category.UserID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *categoryRepository) Delete(userID int64, id int) error {
	query := `DELETE FROM categories WHERE id = ? AND userId = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, userID)
	return err
}
