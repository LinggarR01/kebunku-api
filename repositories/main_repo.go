package repositories

import (
	"database/sql"
	"kebunku-api/config"
	"log"
)

type CompRepository interface{}

type compRepository struct {
	DB *sql.DB
}

func NewComponentRepository(DB *sql.DB) *compRepository {
	db := config.InitDB()

	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS tanaman (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    class VARCHAR(255) NOT NULL,
    plant_order VARCHAR(255) NOT NULL,
    family VARCHAR(255) NOT NULL
);

		`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return &compRepository{
		DB: db,
	}
}
