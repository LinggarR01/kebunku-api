package services

import (
	"kebunku-api/repositories"
)

type CompService interface {

}


type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}