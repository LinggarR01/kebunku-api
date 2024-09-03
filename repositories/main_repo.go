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
    BEGIN TRY
        CREATE TABLE tanaman (
            id INT IDENTITY(1,1) PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
			description VARCHAR(MAX),
			class VARCHAR(255) NOT NULL,
			order VARCHAR(255) NOT NULL,
			family VARCHAR(255) NOT NULL,
        )
		END TRY
		BEGIN CATCH
			-- Ignore the error if the table already exists
		END CATCH
		`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return &compRepository{
		DB: db,
	}
}
