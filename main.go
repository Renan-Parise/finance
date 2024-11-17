package main

import (
	"github.com/Renan-Parise/finances/internal/container"
	"github.com/Renan-Parise/finances/internal/db"
	"github.com/Renan-Parise/finances/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.RunMigrations()

	container := container.NewContainer()

	router := gin.Default()

	handlers.NewTransactionHandler(router, container.TransactionUseCase)
	handlers.NewCategoryHandler(router, container.CategoryUseCase)

	router.Run("127.0.0.1:8180")
}
