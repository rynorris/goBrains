// This package listens to various events and keeps track of statistics relating to the simulation over time.
package stats

import (
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/events"
)

// Contains stats about a creature population.
// Should perhaps be refactored to be more general later, but this is good for now.
type stats struct {
	// Instantaneous statistics.
	population int
	averageAge int

	// Long-term statistics.
	largestAge     int
	oldestCreature *creature.Creature
}

var (
	Global stats
)

// Starts up the global stats collector listening to global events.
func Start() {
	Global.Listen(events.Global)
}

// Listen to events from the given event handler.
func (s *stats) Listen(h events.Handler) {
	h.Register(events.POPULATION_STATE, func(ev events.Event) {
		s.PopulationState(ev.(events.PopulationEvent))
	})
	h.Register(events.ENTITY_DESTROY, func(ev events.Event) {
		s.EntityDestroy(ev.(events.EntityEvent))
	})
}

// Function to update stats on creature death.
func (s *stats) EntityDestroy(ev events.EntityEvent) {
	e := ev.E

	switch e := e.(type) {
	case *creature.Creature:
		if e.Age() > s.largestAge {
			s.largestAge = e.Age()
			s.oldestCreature = e
		}
	}
}

// Update stats on Population State Event.
func (s *stats) PopulationState(ev events.PopulationEvent) {
	s.population = ev.Population
	s.averageAge = ev.AverageAge
}

// Getter Methods
func (s *stats) Population() int {
	return s.population
}

func (s *stats) Oldest() (int, *creature.Creature) {
	return s.largestAge, s.oldestCreature
}

func (s *stats) AverageAge() int {
	return s.averageAge
}
