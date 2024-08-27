package repositories

import (
	"kebunku-api/config"
	"database/sql"
)

type CompRepository interface {}

type compRepository struct {
	DB *sql.DB
}

func NewComponentRepository(DB *sql.DB) *compRepository {
	db := config.InitDB()

	return &compRepository{
		DB: db,
	}
}