package entitymanager

import (
	"math/rand"

	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/food"
)

// spawnFood creates a new blob of food in a random location
// and adds it to the food list.
func (m *em) spawnFood() {
	f := food.New(m.lm, 1000)
	m.food[f] = struct{}{}
}

// breedRandom breeds two random creatures and adds the resulting
// creature to the creature list.
func (m *em) breedRandom() {
	// Do nothing if there are no creatures.
	if len(m.creatures) == 0 {
		return
	}

	// Get two random indices which are not the same.
	ix1 := rand.Intn(len(m.creatures))
	ix2 := ix1
	for ix2 == ix1 && len(m.creatures) > 1 {
		ix2 = rand.Intn(len(m.creatures))
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
	for c, _ := range m.creatures {
		switch ii {
		case ix1:
			mother = c.(*creature.Creature)
		case ix2:
			father = c.(*creature.Creature)
			break
		}
		ii++
	}

	// Breed the creatures.
	child := mother.Breed(father)

	m.creatures[child] = struct{}{}
}
