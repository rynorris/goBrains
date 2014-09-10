/*
 * Entity Manager
 *
 * em is in charge of all entities in the simulation.
 * It handles spawning and updating creatures and food.
 *
 */
package entitymanager

import (
	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/locationmanager"
)

type em struct {
	creatures      entityList
	food           entityList
	lm             locationmanager.LM
	breeding_timer int
	food_timer     int
	stats_timer    int
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
		stats_timer:    0,
	}
}

// Reset initialises a new em instance.
// Spawn a fresh set of random creatures and food.
func (m *em) Reset() {
	// Reset variables.
	m.breeding_timer = 0
	m.food_timer = 0
	m.stats_timer = 0

	// Construct LM.
	m.lm = locationmanager.New()

	// Reset creatures.
	m.creatures.Clear()
	for i := 0; i < config.Global.Entity.InitialCreatures; i++ {
		newCreature := creature.NewSimple(m.lm)
		m.creatures[newCreature] = struct{}{}
	}

	// Reset food
	m.food.Clear()
	for i := 0; i < config.Global.Entity.InitialFood; i++ {
		f := food.New(m.lm, config.Global.Entity.FoodSize)
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
	if m.breeding_timer >= config.Global.Entity.BreedingRate {
		m.breedRandom()
		m.breeding_timer = 0
	}

	// Spawn new food if necessary.
	m.food_timer++
	if m.food_timer >= config.Global.Entity.FoodSpawnRate {
		m.spawnFood()
		m.food_timer = 0
	}

	// Report state of simulation as events if necessary.
	m.stats_timer++
	if m.stats_timer >= 100 {
		m.report()
		m.stats_timer = 0
	}
}

// Entities returns a slice containing all the entities
// in the simulation.
func (m *em) Entities() []entity.Entity {
	return append(m.food.Slice(), m.creatures.Slice()...)
}

// Report state of entities as a global event.
func (m *em) report() {
	averageAge := 0
	ev := events.PopulationEvent{
		Id:         1,
		Population: len(m.creatures),
		AverageAge: averageAge,
	}

	events.Global.Broadcast(ev)
}
