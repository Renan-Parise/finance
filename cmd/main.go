package main

import (
	"github.com/Renan-Parise/finances/internal/db"
	"github.com/Renan-Parise/finances/internal/repositories"
	"github.com/Renan-Parise/finances/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.GetDB()

	transactionRepo := repositories.NewTransactionrepositories(database)

	transactionUseCase := usecases.NewTransactionUseCase(transactionRepo)

	router := gin.Default()

	handlers.NewTransactionHandler(router, transactionUseCase)

	router.Run(":8080")
}
