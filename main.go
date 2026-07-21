package main

import (
	"finance-api/database"
	"finance-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	router.POST("/transactions", handlers.CreateTransaction)
	router.GET("/transactions", handlers.GetTransactions)
	router.DELETE("/transactions/:id", handlers.DeleteTransactionByID)

	router.Run(":8080")
}