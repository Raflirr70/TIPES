package seeders

import (
	"log"

	model "tipes/internal/model"

	"gorm.io/gorm"
)

func SeedAirports(db *gorm.DB) {
	airports := []model.Airport{
		{City: "Jakarta", Country: "Indonesia", AirportCode: "CGK"},
		{City: "Surabaya", Country: "Indonesia", AirportCode: "SUB"},
		{City: "Denpasar", Country: "Indonesia", AirportCode: "DPS"},
		{City: "Medan", Country: "Indonesia", AirportCode: "KNO"},
		{City: "Makassar", Country: "Indonesia", AirportCode: "UPG"},
		{City: "Yogyakarta", Country: "Indonesia", AirportCode: "YIA"},
		{City: "Bandung", Country: "Indonesia", AirportCode: "BDO"},
		{City: "Singapore", Country: "Singapore", AirportCode: "SIN"},
		{City: "Kuala Lumpur", Country: "Malaysia", AirportCode: "KUL"},
		{City: "Bangkok", Country: "Thailand", AirportCode: "BKK"},
	}

	if err := db.Create(&airports).Error; err != nil {
		log.Fatal("Failed to seed airports:", err)
	}

	log.Println("Airport seeder executed successfully")
}
