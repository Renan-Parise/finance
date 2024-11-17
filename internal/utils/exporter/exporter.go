package exporter

import (
	"io"

	"github.com/Renan-Parise/finances/internal/entities"
)

type ExportStrategy interface {
	Export(transactions []*entities.Transaction, w io.Writer) error
}

type ExportContext struct {
	strategy ExportStrategy
}

func (e *ExportContext) SetStrategy(strategy ExportStrategy) {
	e.strategy = strategy
}

func (e *ExportContext) Execute(transactions []*entities.Transaction, w io.Writer) error {
	return e.strategy.Export(transactions, w)
}
