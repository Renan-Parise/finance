package http

import (
	"net/http"
	"strconv"

	"github.com/Renan-Parise/finances/internal/entities"
	"github.com/Renan-Parise/finances/internal/middlewares"
	"github.com/Renan-Parise/finances/internal/usecases"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUseCase usecases.TransactionUseCase
}

func NewTransactionHandler(router *gin.Engine, tu usecase.TransactionUseCase) {
	handler := &TransactionHandler{
		transactionUseCase: tu,
	}

	transactions := router.Group("/transactions")
	transactions.Use(middlewares.JWTAuthMiddleware())
	{
		transactions.POST("/", handler.CreateTransaction)
		transactions.GET("/", handler.GetTransactions)
		transactions.PUT("/:id", handler.UpdateTransaction)
		transactions.DELETE("/:id", handler.DeleteTransaction)
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		Description string  `json:"description" binding:"required"`
		Category    string  `json:"category" binding:"required"`
		Amount      float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.transactionUseCase.CreateTransaction(userID.(int64), input.Description, input.Category, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	transactions, err := h.transactionUseCase.GetTransactions(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var input struct {
		Description string  `json:"description" binding:"required"`
		Category    string  `json:"category" binding:"required"`
		Amount      float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := &entities.Transaction{
		ID:          id,
		UserID:      userID.(int64),
		Description: input.Description,
		Category:    input.Category,
		Amount:      input.Amount,
	}

	err = h.transactionUseCase.UpdateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	err = h.transactionUseCase.DeleteTransaction(userID.(int64), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
