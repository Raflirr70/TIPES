package models

type User struct {
	UserID      uint `gorm:"primaryKey"`
	Name        string
	Email       string
	NoTelephone string
	Password    string
	IsActive    bool
}
