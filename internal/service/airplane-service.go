package service

import (
	model "tipes/internal/model"
	"tipes/internal/repository"
)

type AirplaneService interface {
	GetAll(name string) ([]model.Airplane, error)
}

type airplaneService struct {
	repo repository.AirplaneRepository
}

func NewAirplaneService(repo repository.AirplaneRepository) AirplaneService {
	return &airplaneService{repo}
}

func (s *airplaneService) GetAll(name string) ([]model.Airplane, error) {
	return s.repo.FindWithFilter(name)
}
