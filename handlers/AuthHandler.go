package handlers

import (
	"net/http"

	"finance-api/database"
	"finance-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao processar senha"})
		return
	}

	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err = database.DB.QueryRow(query, newUser.Name, newUser.Email, string(hash)).
		Scan(&newUser.ID, &newUser.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email já cadastrado ou erro ao criar usuário"})
		return
	}

	newUser.Password = ""
	c.JSON(http.StatusCreated, newUser)
}