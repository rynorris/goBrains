/*
 * Entity Manager
 *
 * EM is in charge of all entities in the simulation.
 * It handles spawning and updating Creatures and food.
 *
 */
package entitymanager

import (
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/locationmanager"
	"math/rand"
)

var Creatures map[*creature.Creature]struct{}
var lm locationmanager.Detection
var breeding_timer int

// Start up a new simulation.
// Spawn a fresh set of random creatures and food.
func Start() {
	// Reset variables.
	breeding_timer = 0
	Creatures = map[*creature.Creature]struct{}{}

	// Construct LM.
	lm = locationmanager.NewLocationManager()

	// Reset creatures.
	for c, _ := range Creatures {
		delete(Creatures, c)
	}
	for i := 0; i < initial_creatures; i++ {
		newCreature := creature.NewSimple(lm)
		Creatures[newCreature] = struct{}{}
	}
}

// Perform one update cycle.
// Update all creatures once.
// Spawn new creatures/food if necessary.
func Spin() {
	// Update all creatures
	for c, _ := range Creatures {
		if c.Check() {
			// Uh oh! This creature died! :(
			delete(Creatures, c)
		}
	}

	// Spawn new creatures if necessary.
	breeding_timer++
	if breeding_timer >= breeding_rate {
		breedRandom()
		breeding_timer = 0
	}
}

// Breeds two random creatures and adds the resulting
// creature to the creature list.
func breedRandom() {
	// Get two random indices which are not the same.
	ix1 := rand.Intn(len(Creatures))
	ix2 := ix1
	for ix2 == ix1 {
		ix2 = rand.Intn(len(Creatures))
	}

	// Ensure they are in order.
	// This allows us to break out of the creature loop early.
	if ix2 < ix1 {
		ix1, ix2 = ix2, ix1
	}

	// Resolve indices into creatures.
	// Because of the random nature of iterating over a map,
	// A particular index will not map to the same creature
	// each time. However we don't really care, so long as
	// it's random, and we get 2 different creatures to breed.
	var mother, father *creature.Creature
	ii := 0
	for c, _ := range Creatures {
		switch ii {
		case ix1:
			mother = c
		case ix2:
			father = c
			break
		}
		ii++
	}

	// Breed the creatures.
	child := mother.Breed(father)

	Creatures[child] = struct{}{}
}
