package models

type SeatClass struct {
	SeatClassID uint `gorm:"primaryKey"`
	Name        string
	Price_adult int
	Price_child int
	Price_baby  int
}
