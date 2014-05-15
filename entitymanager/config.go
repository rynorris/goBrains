package entitymanager

import "github.com/DiscoViking/goBrains/config"

var (
	breeding_rate       int
	food_replenish_rate int
	initial_creatures   int
	initial_food        int
	food_size           float64
)

func LoadConfig(cfg *config.Config) {
	initial_creatures = cfg.Entity.InitialCreatures
	initial_food = cfg.Entity.InitialFood
	breeding_rate = cfg.Entity.BreedingRate
	food_replenish_rate = cfg.Entity.FoodSpawnRate
	food_size = cfg.Entity.FoodSize
}
