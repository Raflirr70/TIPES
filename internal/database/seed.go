package database

import (
	config "tipes/internal/config"
	seed "tipes/internal/seed"
)

func Seed() {
	seed.SeedAirports(config.DB)
	seed.SeedAirplanes(config.DB)
	seed.SeedFlight(config.DB)
}
