package models

import (
	"time"
)

type Flight struct {
	FlightID      uint `gorm:"primaryKey"`
	AirportIDFrom uint
	AirportIDTo   uint
	AirplaneID    uint
	SeatClassID   uint
	LeftOver      int
	Date          time.Time

	AirportFrom Airport `gorm:"foreignKey:AirportIDFrom"`
	AirportTo   Airport `gorm:"foreignKey:AirportIDTo"`
	Airplane    Airplane
	SeatClass   SeatClass
}
