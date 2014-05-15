package entitymanager

import "github.com/DiscoViking/goBrains/config"

var (
	breeding_rate       int     = 150  // Spawn a new creature every this many ticks.
	food_replenish_rate int     = 200  // Spawn a new blob of food every this many ticks.
	initial_creatures   int     = 40   // Start with this many random creatures.
	initial_food        int     = 50   // Start with this many randomly placed food blobs.
	food_size           float64 = 1000 // Size of food blobs to spawn.
)

func LoadConfig(cfg *config.Config) {
	initial_creatures = cfg.EntityManager.InitialCreatures
	initial_food = cfg.EntityManager.InitialFood
	breeding_rate = cfg.EntityManager.BreedingRate
	food_replenish_rate = cfg.EntityManager.FoodSpawnRate
	food_size = cfg.EntityManager.FoodSize
}
