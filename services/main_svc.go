package services

import (
	"kebunku-api/dto"
	"kebunku-api/repositories"
)

type CompService interface {
	RegisterTanaman(data dto.Tanaman) error
	GetTanaman() ([]dto.Tanaman, error)
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
