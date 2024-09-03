package services

import "kebunku-api/dto"

func (s *compServices) RegisterTanaman(data dto.Tanaman) error {
	return s.repo.RegisterTanaman(data)
}
func (s *compServices) GetTanaman() ([]dto.Tanaman, error) {
	return s.repo.GetTanaman()
}

func (s *compServices) UploadTanaman(file_url string, id string) error {
	return s.repo.UploadTanaman(file_url, id)
}
