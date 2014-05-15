package config

// Config holds all the configuration data for goBrains
type Config struct {
	General struct {
		ScreenWidth  int // Width of screen
		ScreenHeight int // Height of screen
	}

	EntityManager struct {
		InitialCreatures int     // The number of creatures to start with.
		InitialFood      int     // The number of food blods to start with.
		BreedingRate     int     // Spawn a new creatures every this many ticks.
		FoodSpawnRate    int     // Spawn a new food blob every this many ticks.
		FoodSize         float64 // Size of food blobs to spawn.
	}
}
