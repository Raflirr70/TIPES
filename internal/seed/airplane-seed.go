package seeders

import (
	"fmt"
	"log"

	model "tipes/internal/model"
	models "tipes/internal/model"

	"gorm.io/gorm"
)

func GenerateSeatsFromTotal(totalSeat int) []models.Seat {
	var seats []models.Seat

	seatsPerRow := 6 // A–F
	seatCount := 0

	for row := 1; seatCount < totalSeat; row++ {
		for col := 0; col < seatsPerRow && seatCount < totalSeat; col++ {

			seatNumber := fmt.Sprintf("%d%c", row, 'A'+col)

			seats = append(seats, models.Seat{
				NoSeat: seatNumber,
				Status: false, // available
			})

			seatCount++
		}
	}

	return seats
}

func SeedAirplanes(db *gorm.DB) {
	airplanes := []model.Airplane{
		{Name: "Boeing 737-800", NSeat: 189, Baggage: 20, CabinBaggage: 7},
		{Name: "Boeing 737-900", NSeat: 220, Baggage: 25, CabinBaggage: 7},
		{Name: "Airbus A320", NSeat: 180, Baggage: 20, CabinBaggage: 7},
		{Name: "Airbus A321", NSeat: 220, Baggage: 25, CabinBaggage: 7},
		{Name: "Boeing 747-400", NSeat: 416, Baggage: 30, CabinBaggage: 10},
		{Name: "Boeing 787 Dreamliner", NSeat: 290, Baggage: 30, CabinBaggage: 10},
		{Name: "Airbus A330", NSeat: 250, Baggage: 30, CabinBaggage: 10},
		{Name: "Airbus A350", NSeat: 300, Baggage: 35, CabinBaggage: 10},
		{Name: "ATR 72", NSeat: 70, Baggage: 15, CabinBaggage: 5},
		{Name: "Bombardier CRJ-1000", NSeat: 100, Baggage: 20, CabinBaggage: 7},

		{Name: "Boeing 737-700", NSeat: 149, Baggage: 20, CabinBaggage: 7},
		{Name: "Boeing 737 MAX 8", NSeat: 178, Baggage: 25, CabinBaggage: 7},
		{Name: "Boeing 737 MAX 9", NSeat: 220, Baggage: 25, CabinBaggage: 7},
		{Name: "Airbus A220-100", NSeat: 120, Baggage: 20, CabinBaggage: 7},
		{Name: "Airbus A220-300", NSeat: 145, Baggage: 20, CabinBaggage: 7},
		{Name: "Embraer E175", NSeat: 88, Baggage: 15, CabinBaggage: 5},
		{Name: "Embraer E190", NSeat: 114, Baggage: 20, CabinBaggage: 7},
		{Name: "Embraer E195", NSeat: 132, Baggage: 20, CabinBaggage: 7},
		{Name: "Boeing 777-200", NSeat: 305, Baggage: 30, CabinBaggage: 10},
		{Name: "Boeing 777-300ER", NSeat: 396, Baggage: 35, CabinBaggage: 10},

		{Name: "Airbus A380", NSeat: 525, Baggage: 40, CabinBaggage: 15},
		{Name: "Boeing 767-300", NSeat: 218, Baggage: 25, CabinBaggage: 7},
		{Name: "Boeing 757-200", NSeat: 239, Baggage: 25, CabinBaggage: 7},
		{Name: "Airbus A319", NSeat: 140, Baggage: 20, CabinBaggage: 7},
		{Name: "Airbus A318", NSeat: 107, Baggage: 20, CabinBaggage: 7},
		{Name: "Comac C919", NSeat: 158, Baggage: 20, CabinBaggage: 7},
		{Name: "Irkut MC-21", NSeat: 165, Baggage: 20, CabinBaggage: 7},
		{Name: "Sukhoi Superjet 100", NSeat: 108, Baggage: 20, CabinBaggage: 7},
		{Name: "Antonov An-148", NSeat: 85, Baggage: 15, CabinBaggage: 5},
		{Name: "Tupolev Tu-204", NSeat: 210, Baggage: 25, CabinBaggage: 7},

		{Name: "Boeing 737-500", NSeat: 132, Baggage: 20, CabinBaggage: 7},
		{Name: "Boeing 737-600", NSeat: 108, Baggage: 20, CabinBaggage: 7},
		{Name: "Airbus A340", NSeat: 261, Baggage: 30, CabinBaggage: 10},
		{Name: "Boeing 787-9", NSeat: 290, Baggage: 30, CabinBaggage: 10},
		{Name: "Boeing 787-10", NSeat: 330, Baggage: 35, CabinBaggage: 10},
		{Name: "Airbus A321neo", NSeat: 220, Baggage: 25, CabinBaggage: 7},
		{Name: "Airbus A320neo", NSeat: 180, Baggage: 20, CabinBaggage: 7},
		{Name: "ATR 42", NSeat: 48, Baggage: 15, CabinBaggage: 5},
		{Name: "Dash 8 Q400", NSeat: 78, Baggage: 15, CabinBaggage: 5},

		{Name: "Embraer E170", NSeat: 72, Baggage: 15, CabinBaggage: 5},
		{Name: "Embraer E175-E2", NSeat: 88, Baggage: 20, CabinBaggage: 7},
		{Name: "Embraer E190-E2", NSeat: 114, Baggage: 20, CabinBaggage: 7},
		{Name: "Embraer E195-E2", NSeat: 132, Baggage: 25, CabinBaggage: 7},
		{Name: "Boeing 777-8", NSeat: 350, Baggage: 35, CabinBaggage: 10},
	}

	for _, p := range airplanes {

		// Prevent duplicate seed
		// var existing model.Airplane
		// if err := db.Where("airplane_id = ?", p.AirplaneID).First(&existing).Error; err == nil {
		// 	continue
		// }

		// Generate seats based on total seat
		p.Seat = GenerateSeatsFromTotal(p.NSeat)

		if err := db.Create(&p).Error; err != nil {
			log.Fatal("Failed to seed seat:", err)
		}
	}

	log.Println("Airplane seeder (50 data) executed successfully")
}
