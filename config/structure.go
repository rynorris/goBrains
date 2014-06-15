package config

// Config holds all the configuration data for goBrains
type Config struct {
	General struct {
		ScreenWidth  int // Width of screen
		ScreenHeight int // Height of screen
	}

	Entity struct {
		InitialCreatures int     // The number of creatures to start with.
		InitialFood      int     // The number of food blods to start with.
		BreedingChance   int     // Chance for each creature to breed each tick is 1/n.
		FoodSpawnRate    int     // Spawn a new food blob every this many ticks.
		FoodSize         float64 // Size of food blobs to spawn.
	}

	Brain struct {
		ChargeDecayRate     float64
		SynapseMaxCharge    float64
		SynapseOutputScale  float64
		NodeFiringThreshold float64
		NodeFiringStrength  float64
	}

	Genetics struct {
		MutationRate int
	}
}
