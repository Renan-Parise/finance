package exporter

import (
	"fmt"
	"io"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/xuri/excelize/v2"
)

type XLSXExporter struct{}

func (e *XLSXExporter) Export(transactions []*entities.Transaction, w io.Writer) error {
	file := excelize.NewFile()
	sheet := "Transactions"
	file.NewSheet(sheet)
	file.SetCellValue(sheet, "A1", "ID")
	file.SetCellValue(sheet, "B1", "Description")
	file.SetCellValue(sheet, "C1", "Category")
	file.SetCellValue(sheet, "D1", "Amount")
	file.SetCellValue(sheet, "E1", "Created At")
	file.SetCellValue(sheet, "F1", "Updated At")

	for i, t := range transactions {
		row := i + 2
		file.SetCellValue(sheet, fmt.Sprintf("A%d", row), t.ID)
		file.SetCellValue(sheet, fmt.Sprintf("B%d", row), t.Description)
		file.SetCellValue(sheet, fmt.Sprintf("C%d", row), t.Category)
		file.SetCellValue(sheet, fmt.Sprintf("D%d", row), t.Amount)
		file.SetCellValue(sheet, fmt.Sprintf("E%d", row), t.CreatedAt)
		file.SetCellValue(sheet, fmt.Sprintf("F%d", row), t.UpdatedAt)
	}

	return file.Write(w)
}
