package models

import (
	"time"
)

type Order struct {
	OrderID     uint `gorm:"primaryKey"`
	UserID      uint
	FlightID    uint
	FullName    string
	FamilyName  string
	NoTelephone int
	Email       time.Time
	TotalPrice  int

	User   User
	Flight Flight

	Ticket []Ticket
}
