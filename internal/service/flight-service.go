package service

import (
	model "tipes/internal/model"
	"tipes/internal/repository"
)

type FlightService interface {
	GetAll(from, to, airplaneName, seatClassName string) ([]model.Flight, error)
}

type flightService struct {
	repo repository.FlightRepository
}

func NewFlightService(repo repository.FlightRepository) FlightService {
	return &flightService{repo}
}

func (s *flightService) GetAll(from, to, airplaneName, seatClassName string) ([]model.Flight, error) {
	return s.repo.FindWithFilter(from, to, airplaneName, seatClassName)
}
