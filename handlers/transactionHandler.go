package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

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

func GetTransactions(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT id, description, amount, type, category, created_at
		FROM transactions
		ORDER BY created_at DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	
	var transactions []models.Transaction

	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.Description, &t.Amount, &t.Type, &t.Category, &t.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, t)
	}

	c.JSON(http.StatusOK, transactions)
}

func DeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")

	result, err := database.DB.Exec("DELETE FROM transactions WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

func UpdateTransactionByID(c *gin.Context) {
	id := c.Param("id")

	var updatedTransaction models.Transaction
	if err := c.ShouldBindJSON(&updatedTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE transactions
		SET description = $1, amount = $2, type = $3, category = $4
		WHERE id = $5
		RETURNING id, description, amount, type, category, created_at
	`

	err := database.DB.QueryRow(
		query,
		updatedTransaction.Description,
		updatedTransaction.Amount,
		updatedTransaction.Type,
		updatedTransaction.Category,
		id,
	).Scan(
		&updatedTransaction.ID,
		&updatedTransaction.Description,
		&updatedTransaction.Amount,
		&updatedTransaction.Type,
		&updatedTransaction.Category,
		&updatedTransaction.CreatedAt,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTransaction)
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var transaction models.Transaction
	err = database.DB.QueryRow(`
		SELECT id, description, amount, type, category, created_at
		FROM transactions
		WHERE id = $1
	`, idInt).Scan(&transaction.ID, &transaction.Description, &transaction.Amount, &transaction.Type, &transaction.Category, &transaction.CreatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}