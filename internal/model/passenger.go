package models

import (
	"time"
)

type Passenger struct {
	PassengerID   uint `gorm:"primaryKey"`
	TicketID      uint
	Title         string
	FullName      string
	FamilyName    string
	BirthDate     time.Time
	Pasport       string
	OriginCountry string
	Expired       time.Time

	Ticket Ticket
}
