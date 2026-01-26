package repository

import (
	model "tipes/internal/model"

	"gorm.io/gorm"
)

type FlightRepository interface {
	FindWithFilter(from, to, airplaneName, seatClassName string) ([]model.Flight, error)
}

type flightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) FlightRepository {
	return &flightRepository{db}
}

func (r *flightRepository) FindWithFilter(from, to, airplaneName, seatClassName string) ([]model.Flight, error) {
	var flights []model.Flight

	query := r.db.Model(&model.Flight{}).
		Joins("JOIN airports as city_from ON city_from.airport_id = flights.airport_id_from").
		Joins("JOIN airports as city_to ON city_to.airport_id = flights.airport_id_to").
		Joins("JOIN airplanes as ap ON ap.airplane_id = flights.airplane_id").
		Joins("JOIN seat_classes as sc ON sc.seat_class_id = flights.seat_class_id")

	if from != "" {
		query = query.Where("city_from.city ILIKE ?", "%"+from+"%")
	}
	if to != "" {
		query = query.Where("city_to.city ILIKE ?", "%"+to+"%")
	}
	if airplaneName != "" {
		query = query.Where("ap.name ILIKE ?", "%"+airplaneName+"%")
	}
	if seatClassName != "" {
		query = query.Where("sc.name ILIKE ?", "%"+seatClassName+"%")
	}

	err := query.Preload("AirportFrom").Preload("AirportTo").Preload("Airplane").Preload("SeatClass").Find(&flights).Error
	return flights, err
}
