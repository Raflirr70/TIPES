package models

type Payment struct {
	PaymentID uint `gorm:"primaryKey"`
	OrderID   uint
	Method    string
	Amount    int
	status    bool

	Order Order
}
