package models

type Airplane struct {
	AirplaneID   uint `gorm:"primaryKey"`
	Name         string
	NSeat        int
	Baggage      int
	CabinBaggage int
}
