package database

import (
	"log"
	"tipes/internal/config"
	model "tipes/internal/model"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&model.User{},
		&model.Airport{},
		&model.SeatClass{},
		&model.Airplane{},
		&model.Seat{},
		&model.Flight{},
		&model.Order{},
		&model.Payment{},
		&model.Passenger{},
		&model.Ticket{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database migrated")
}
