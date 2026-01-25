package repository

import (
	model "tipes/internal/model"

	"gorm.io/gorm"
)

type AirplaneRepository interface {
	FindWithFilter(name string) ([]model.Airplane, error)
}

type airplaneRepository struct {
	db *gorm.DB
}

func NewAirplaneRepository(db *gorm.DB) AirplaneRepository {
	return &airplaneRepository{db}
}

func (r *airplaneRepository) FindWithFilter(name string) ([]model.Airplane, error) {
	var airplanes []model.Airplane

	query := r.db.Model(&model.Airplane{})

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Find(&airplanes).Error
	return airplanes, err
}
