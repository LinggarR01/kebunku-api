package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func InitDB() *sql.DB {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	dbConfig := DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	dbURI := fmt.Sprintf("%s:%s@/kebunku", dbConfig.User, dbConfig.Password)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
