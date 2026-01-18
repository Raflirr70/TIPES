package models

type Airport struct {
	AirportID   uint `gorm:"primaryKey"`
	City        string
	Country     string
	AirportCode string
}
