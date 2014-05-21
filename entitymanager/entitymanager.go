/*
 * Entity Manager
 *
 * em is in charge of all entities in the simulation.
 * It handles spawning and updating creatures and food.
 *
 */
package entitymanager

import (
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/locationmanager"
)

type em struct {
	creatures      entityList
	food           entityList
	lm             locationmanager.LM
	breeding_timer int
	food_timer     int
}

func (m *em) LocationManager() locationmanager.LM {
	return m.lm
}

// New returns a new em instance.
func New() Manager {
	return &em{
		creatures:      map[entity.Entity]struct{}{},
		food:           map[entity.Entity]struct{}{},
		lm:             locationmanager.New(),
		breeding_timer: 0,
		food_timer:     0,
	}
}

// Reset initialises a new em instance.
// Spawn a fresh set of random creatures and food.
func (m *em) Reset() {
	// Reset variables.
	m.breeding_timer = 0

	// Construct LM.
	m.lm = locationmanager.New()

	// Reset creatures.
	for c, _ := range m.creatures {
		delete(m.creatures, c)
	}
	for i := 0; i < initial_creatures; i++ {
		newCreature := creature.NewSimple(m.lm)
		m.creatures[newCreature] = struct{}{}
	}

	// Reset food
	for f, _ := range m.food {
		delete(m.food, f)
	}
	for i := 0; i < initial_food; i++ {
		f := food.New(m.lm, food_size)
		m.food[f] = struct{}{}
	}
}

// Spin performs one update cycle.
// Update all creatures once.
// Spawn new creatures/food if necessary.
func (m *em) Spin() {
	// Update all creatures
	m.creatures.Work()
	m.creatures.Check()

	// Update all food
	m.food.Check()

	// Spawn new creatures if necessary.
	m.breeding_timer++
	if m.breeding_timer >= breeding_rate {
		m.breedRandom()
		m.breeding_timer = 0
	}

	// Spawn new food if necessary.
	m.food_timer++
	if m.food_timer >= food_replenish_rate {
		m.spawnFood()
		m.food_timer = 0
	}
}

// Entities returns a slice containing all the entities
// in the simulation.
func (m *em) Entities() []entity.Entity {
	return append(m.food.Slice(), m.creatures.Slice()...)
}
