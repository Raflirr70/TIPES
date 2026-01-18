package models

type Seat struct {
	SeatID     uint `gorm:"primaryKey"`
	AirplaneID uint
	NoSeat     int
	Status     bool

	Airplane Airplane
}
