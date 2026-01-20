package seeders

import (
	"log"
	"time"

	model "tipes/internal/model"

	"gorm.io/gorm"
)

func SeedFlight(db *gorm.DB) {

	flights := []model.Flight{
		{
			AirportIDFrom: 1, AirportIDTo: 2, AirplaneID: 1, LeftOver: 189,
			Date:      time.Date(2026, 2, 1, 7, 20, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Economy", Price_adult: 750000, Price_child: 500000, Price_baby: 0},
		},
		{
			AirportIDFrom: 2, AirportIDTo: 1, AirplaneID: 1, LeftOver: 189,
			Date:      time.Date(2026, 2, 1, 10, 30, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Economy", Price_adult: 750000, Price_child: 500000, Price_baby: 0},
		},
		{
			AirportIDFrom: 1, AirportIDTo: 3, AirplaneID: 3, LeftOver: 180,
			Date:      time.Date(2026, 2, 2, 8, 45, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Economy", Price_adult: 680000, Price_child: 450000, Price_baby: 0},
		},
		{
			AirportIDFrom: 3, AirportIDTo: 1, AirplaneID: 3, LeftOver: 180,
			Date:      time.Date(2026, 2, 2, 12, 10, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Economy", Price_adult: 680000, Price_child: 450000, Price_baby: 0},
		},
		{
			AirportIDFrom: 1, AirportIDTo: 4, AirplaneID: 20, LeftOver: 396,
			Date:      time.Date(2026, 2, 3, 13, 10, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Business", Price_adult: 3500000, Price_child: 2500000, Price_baby: 500000},
		},
		{
			AirportIDFrom: 4, AirportIDTo: 1, AirplaneID: 20, LeftOver: 396,
			Date:      time.Date(2026, 2, 3, 18, 30, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Business", Price_adult: 3500000, Price_child: 2500000, Price_baby: 500000},
		},
		{
			AirportIDFrom: 3, AirportIDTo: 5, AirplaneID: 9, LeftOver: 70,
			Date:      time.Date(2026, 2, 4, 6, 30, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Economy", Price_adult: 350000, Price_child: 250000, Price_baby: 0},
		},
		{
			AirportIDFrom: 5, AirportIDTo: 3, AirplaneID: 9, LeftOver: 70,
			Date:      time.Date(2026, 2, 4, 9, 0, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "Economy", Price_adult: 350000, Price_child: 250000, Price_baby: 0},
		},
		{
			AirportIDFrom: 1, AirportIDTo: 6, AirplaneID: 21, LeftOver: 525,
			Date:      time.Date(2026, 2, 5, 22, 0, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "First Class", Price_adult: 8500000, Price_child: 6000000, Price_baby: 1500000},
		},
		{
			AirportIDFrom: 6, AirportIDTo: 1, AirplaneID: 21, LeftOver: 525,
			Date:      time.Date(2026, 2, 6, 5, 45, 0, 0, time.UTC),
			SeatClass: model.SeatClass{Name: "First Class", Price_adult: 8500000, Price_child: 6000000, Price_baby: 1500000},
		},

		// 20 flight tambahan (pola konsisten)
	}

	// Tambahkan otomatis hingga total 30 flight
	for i := 0; i < 20; i++ {
		flights = append(flights, model.Flight{
			AirportIDFrom: uint((i % 6) + 1),
			AirportIDTo:   uint(((i + 2) % 6) + 1),
			AirplaneID:    uint((i % 10) + 1),
			LeftOver:      150,
			Date:          time.Date(2026, 2, 7+i, 8+(i%10), 0, 0, 0, time.UTC),
			SeatClass: model.SeatClass{
				Name:        "Economy",
				Price_adult: 600000 + (i * 10000),
				Price_child: 400000 + (i * 8000),
				Price_baby:  0,
			},
		})
	}

	if err := db.Create(&flights).Error; err != nil {
		log.Fatal("Failed to seed Flight:", err)
	}

	log.Println("Flight seeder (30 data) executed successfully")
}
