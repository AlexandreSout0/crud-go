package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // driver banco
)

// Abre conexão com o banco de dados
func InitDB() (*sql.DB, error) {

	connectionString := "postgres://postgres:123@localhost:5432/db?sslmode=disable"

	db, erro := sql.Open("postgres", connectionString)
	if erro != nil {
		log.Fatalf("Error opening database: %v", erro)
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		log.Fatalf("Error connecting to database: %v", erro)
		return nil, erro
	}
	return db, nil
}
