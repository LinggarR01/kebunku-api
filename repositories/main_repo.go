package repositories

import (
	"database/sql"
	"kebunku-api/config"
	"kebunku-api/dto"
	"log"
)

type CompRepository interface {
	RegisterTanaman(data dto.Tanaman) error
	GetTanaman() ([]dto.Tanaman, error)
	UploadTanaman (file_url string, id string) error
}

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

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS tanaman_image (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_tanaman INT NOT NULL,
	file_url VARCHAR(255)
);

		`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return &compRepository{
		DB: db,
	}
}
