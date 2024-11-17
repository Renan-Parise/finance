package repositories

import (
	"database/sql"

	"github.com/Renan-Parise/finances/internal/entities"
)

type StatisticsRepository interface {
	GetTotalIncome(userID int64) (float64, error)
	GetTotalExpenses(userID int64) (float64, error)
	GetMostUsedCategory(userID int64) (string, error)
	GetMonthlyExpenses(userID int64) ([]*entities.MonthlyAmount, error)
	GetMonthlyIncome(userID int64) ([]*entities.MonthlyAmount, error)
	GetCategoryMonthlyTotals(userID int64, month, year int) (map[string]float64, error)
}

type statisticsRepository struct {
	db *sql.DB
}

func NewStatisticsRepository(db *sql.DB) StatisticsRepository {
	return &statisticsRepository{db: db}
}

func (r *statisticsRepository) GetTotalIncome(userID int64) (float64, error) {
	query := `SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE userId = ? AND amount > 0`
	var totalIncome float64
	err := r.db.QueryRow(query, userID).Scan(&totalIncome)
	return totalIncome, err
}

func (r *statisticsRepository) GetTotalExpenses(userID int64) (float64, error) {
	query := `SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE userId = ? AND amount < 0`
	var totalExpenses float64
	err := r.db.QueryRow(query, userID).Scan(&totalExpenses)
	return totalExpenses, err
}

func (r *statisticsRepository) GetMostUsedCategory(userID int64) (string, error) {
	query := `
		SELECT c.name, COUNT(*) AS usage_count
		FROM transactions t
		JOIN categories c ON t.category = c.id
		WHERE t.userId = ?
		GROUP BY t.category
		ORDER BY usage_count DESC
		LIMIT 1
	`
	var categoryName string
	var usageCount int

	err := r.db.QueryRow(query, userID).Scan(&categoryName, &usageCount)
	return categoryName, err
}

func (r *statisticsRepository) GetMonthlyExpenses(userID int64) ([]*entities.MonthlyAmount, error) {
	query := `
		SELECT YEAR(createdAt) as year, MONTH(createdAt) as month, ABS(SUM(amount)) as total
		FROM transactions
		WHERE userId = ? AND amount < 0
		GROUP BY YEAR(createdAt), MONTH(createdAt)
		ORDER BY total DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*entities.MonthlyAmount
	for rows.Next() {
		var ma entities.MonthlyAmount
		err := rows.Scan(&ma.Year, &ma.Month, &ma.Total)
		if err != nil {
			return nil, err
		}
		results = append(results, &ma)
	}
	return results, nil
}

func (r *statisticsRepository) GetMonthlyIncome(userID int64) ([]*entities.MonthlyAmount, error) {
	query := `
		SELECT YEAR(createdAt) as year, MONTH(createdAt) as month, SUM(amount) as total
		FROM transactions
		WHERE userId = ? AND amount > 0
		GROUP BY YEAR(createdAt), MONTH(createdAt)
		ORDER BY total DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*entities.MonthlyAmount
	for rows.Next() {
		var ma entities.MonthlyAmount
		err := rows.Scan(&ma.Year, &ma.Month, &ma.Total)
		if err != nil {
			return nil, err
		}
		results = append(results, &ma)
	}
	return results, nil
}

func (r *statisticsRepository) GetCategoryMonthlyTotals(userID int64, month, year int) (map[string]float64, error) {
	query := `
		SELECT c.name, SUM(t.amount) as total
		FROM transactions t
		JOIN categories c ON t.category = c.id
		WHERE t.userId = ? AND MONTH(t.createdAt) = ? AND YEAR(t.createdAt) = ?
		GROUP BY c.name
	`
	rows, err := r.db.Query(query, userID, month, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totals := make(map[string]float64)
	for rows.Next() {
		var categoryName string
		var total float64
		err := rows.Scan(&categoryName, &total)
		if err != nil {
			return nil, err
		}
		totals[categoryName] = total
	}
	return totals, nil
}
