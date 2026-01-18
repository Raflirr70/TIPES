package models

type Ticket struct {
	TicketID      uint `gorm:"primaryKey"`
	OrderID       uint
	SeatID        uint
	TypePassenger string

	Order Order
	Seat  Seat
}
