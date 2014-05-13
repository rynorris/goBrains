package entitymanager

const (
	max_creatures       = 50  // Maximum creatures in play at once. Stop breeding if this many creatures live.
	max_food            = 50  // Maximum number of food blobs in play at once.
	breeding_rate       = 150 // Spawn a new creature every this many ticks.
	food_replenish_rate = 200 // Spawn a new blob of food every this many ticks.
	initial_creatures   = 20  // Start with this many random creatures.
	initial_food        = 20  // Start with this many randomly placed food blobs.
)
