package services

import (
	"kebunku-api/dto"
	"kebunku-api/repositories"
)

type CompService interface {
	RegisterTanaman(data dto.Tanaman) error
	GetTanaman() ([]dto.Tanaman, error)
	UploadTanaman(file_url string, id string) error
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
