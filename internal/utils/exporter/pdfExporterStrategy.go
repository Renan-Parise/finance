package exporter

import (
	"fmt"
	"io"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/jung-kurt/gofpdf"
)

type PDFExporter struct{}

func (e *PDFExporter) Export(transactions []*entities.Transaction, w io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "ID")
	pdf.Cell(50, 10, "Description")
	pdf.Cell(40, 10, "Category")
	pdf.Cell(30, 10, "Amount")
	pdf.Cell(40, 10, "Created At")
	pdf.Ln(10)

	for _, t := range transactions {
		pdf.Cell(40, 10, fmt.Sprintf("%d", t.ID))
		pdf.Cell(50, 10, t.Description)
		pdf.Cell(40, 10, fmt.Sprintf("%d", t.Category))
		pdf.Cell(30, 10, fmt.Sprintf("%.2f", t.Amount))
		pdf.Cell(40, 10, t.CreatedAt.Format("2006-01-02"))
		pdf.Ln(10)
	}

	return pdf.Output(w)
}
