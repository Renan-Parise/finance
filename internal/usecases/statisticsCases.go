package usecases

import (
	"math"
	"time"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/Renan-Parise/finances/internal/repositories"
	"github.com/Renan-Parise/finances/internal/utils"
)

type StatisticsUseCase interface {
	GetCategoryPercentageChanges(userID int64) ([]*entities.CategoryPercentageChange, error)
	GetExpensesByCategory(userID int64) ([]*entities.ExpenseCategorySummary, error)
	GetMonthlyExpensesSummary(userID int64) ([]*entities.MonthlyAmount, error)
	GetGeneralStatistics(userID int64) (*entities.GeneralStatistics, error)
	GetHighestExpenseMonth(userID int64) (*entities.MonthlyAmount, error)
	GetHighestIncomeMonth(userID int64) (*entities.MonthlyAmount, error)
	GetSpendingHeatmap(userID int64) (map[string]float64, error)
}

type statisticsUseCase struct {
	statisticsRepo repositories.StatisticsRepository
}

func NewStatisticsUseCase(sr repositories.StatisticsRepository) StatisticsUseCase {
	return &statisticsUseCase{statisticsRepo: sr}
}

func (uc *statisticsUseCase) GetGeneralStatistics(userID int64) (*entities.GeneralStatistics, error) {
	totalIncome, err := uc.statisticsRepo.GetTotalIncome(userID)
	if err != nil {
		return nil, err
	}
	totalExpenses, err := uc.statisticsRepo.GetTotalExpenses(userID)
	if err != nil {
		return nil, err
	}
	balance := totalIncome + totalExpenses
	mostUsedCategory, err := uc.statisticsRepo.GetMostUsedCategory(userID)
	if err != nil {
		return nil, err
	}
	return &entities.GeneralStatistics{
		TotalIncome:      totalIncome,
		TotalExpenses:    totalExpenses,
		Balance:          balance,
		MostUsedCategory: mostUsedCategory,
	}, nil
}

func (uc *statisticsUseCase) GetHighestExpenseMonth(userID int64) (*entities.MonthlyAmount, error) {
	expenses, err := uc.statisticsRepo.GetMonthlyExpenses(userID)
	if err != nil {
		return nil, err
	}
	if len(expenses) == 0 {
		return nil, nil
	}
	return expenses[0], nil
}

func (uc *statisticsUseCase) GetHighestIncomeMonth(userID int64) (*entities.MonthlyAmount, error) {
	income, err := uc.statisticsRepo.GetMonthlyIncome(userID)
	if err != nil {
		return nil, err
	}
	if len(income) == 0 {
		return nil, nil
	}
	return income[0], nil
}

func (uc *statisticsUseCase) GetCategoryPercentageChanges(userID int64) ([]*entities.CategoryPercentageChange, error) {
	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()
	previousMonth := currentMonth - 1
	previousYear := currentYear
	if previousMonth == 0 {
		previousMonth = 12
		previousYear -= 1
	}

	currentTotals, err := uc.statisticsRepo.GetCategoryMonthlyTotals(userID, int(currentMonth), currentYear)
	if err != nil {
		return nil, err
	}
	previousTotals, err := uc.statisticsRepo.GetCategoryMonthlyTotals(userID, int(previousMonth), previousYear)
	if err != nil {
		return nil, err
	}

	var changes []*entities.CategoryPercentageChange
	for name := range utils.MergeKeys(currentTotals, previousTotals) {
		currentValue := currentTotals[name]
		previousValue := previousTotals[name]
		var percentageChange float64
		var increase bool

		if previousValue != 0 {
			change := currentValue - previousValue
			percentageChange = (change / math.Abs(previousValue)) * 100
			increase = change > 0
		} else {
			percentageChange = 100.0
			increase = currentValue > 0
		}

		changes = append(changes, &entities.CategoryPercentageChange{
			CategoryName:     name,
			PreviousValue:    previousValue,
			CurrentValue:     currentValue,
			PercentageChange: percentageChange,
			Increase:         increase,
		})
	}

	return changes, nil
}

func (uc *statisticsUseCase) GetSpendingHeatmap(userID int64) (map[string]float64, error) {
	return uc.statisticsRepo.GetSpendingHeatmap(userID)
}

func (uc *statisticsUseCase) GetMonthlyExpensesSummary(userID int64) ([]*entities.MonthlyAmount, error) {
	return uc.statisticsRepo.GetMonthlyExpenses(userID)
}

func (uc *statisticsUseCase) GetExpensesByCategory(userID int64) ([]*entities.ExpenseCategorySummary, error) {
	return uc.statisticsRepo.GetExpensesByCategory(userID)
}
