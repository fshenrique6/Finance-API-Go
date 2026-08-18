package main

import (
	"finance-api/database"
	"finance-api/handlers"
	"finance-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	// public routes - dont require authentication
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// protected routes - require authentication
	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.POST("/transactions", handlers.CreateTransaction)
		authorized.GET("/transactions", handlers.GetTransactions)
		authorized.GET("/transactions/:id", handlers.GetTransactionByID)
		authorized.PUT("/transactions/:id", handlers.UpdateTransactionByID)
		authorized.DELETE("/transactions/:id", handlers.DeleteTransactionByID)
		authorized.GET("/summary", handlers.GetSummary)
	}

	router.Run(":8080")
}