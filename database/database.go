package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
	connStr := "postgres://financeapi:financeapi123@localhost:5432/financedb?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("erro ao pingar o banco:", err)
	}

	DB = db
	log.Println("conectado ao banco com sucesso")
}