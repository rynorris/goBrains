package entitymanager

const (
	max_creatures       = 80   // Maximum creatures in play at once. Stop breeding if this many creatures live.
	max_food            = 80   // Maximum number of food blobs in play at once.
	breeding_rate       = 50   // Spawn a new creature every this many ticks.
	food_replenish_rate = 200  // Spawn a new blob of food every this many ticks.
	initial_creatures   = 80   // Start with this many random creatures.
	initial_food        = 50   // Start with this many randomly placed food blobs.
	food_size           = 1000 // Size of food blobs to spawn.
)
