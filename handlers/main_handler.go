package handlers

import "kebunku-api/services"

type compHandlers struct {
	service services.CompService
}

func NewCompHandlers(s services.CompService) *compHandlers {
	return &compHandlers{
		service: s,
	}
}