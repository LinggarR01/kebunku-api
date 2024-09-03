package services

import "kebunku-api/dto"

func (s *compServices) RegisterTanaman(data dto.Tanaman) error {
	return s.repo.RegisterTanaman(data)
}