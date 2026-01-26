package repository

import (
	model "tipes/internal/model"

	"gorm.io/gorm"
)

type AirplaneRepository interface {
	FindAll() ([]model.Airplane, error)
}

type airplaneRepository struct {
	db *gorm.DB
}

func NewAirplaneRepository(db *gorm.DB) AirplaneRepository {
	return &airplaneRepository{db}
}

func (r *airplaneRepository) FindAll() ([]model.Airplane, error) {
	var airplanes []model.Airplane
	err := r.db.Find(&airplanes).Error
	return airplanes, err
}
