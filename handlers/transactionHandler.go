package handlers

import (
	"net/http"

	"finance-api/database"
	"finance-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var newTransaction models.Transaction

	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO transactions (description, amount, type, category)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	err := database.DB.QueryRow(
		query,
		newTransaction.Description,
		newTransaction.Amount,
		newTransaction.Type,
		newTransaction.Category,
	).Scan(&newTransaction.ID, &newTransaction.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTransaction)
}