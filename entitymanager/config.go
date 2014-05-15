package entitymanager

import "github.com/DiscoViking/goBrains/config"

var (
	breeding_rate       int     = 150
	food_replenish_rate int     = 200
	initial_creatures   int     = 40
	initial_food        int     = 50
	food_size           float64 = 1000
)

func LoadConfig(cfg *config.Config) {
	initial_creatures = cfg.Entity.InitialCreatures
	initial_food = cfg.Entity.InitialFood
	breeding_rate = cfg.Entity.BreedingRate
	food_replenish_rate = cfg.Entity.FoodSpawnRate
	food_size = cfg.Entity.FoodSize
}
