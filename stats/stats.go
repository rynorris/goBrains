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
	global stats
)

// Function to update stats based on an Event.
func (s *stats) Update(ev events.Event) {
	switch ev.GetType() {
	case events.POPULATION_STATE:
		s.PopulationState(ev.(events.PopulationEvent))
	case events.ENTITY_DESTROY:
		s.EntityDestroy(ev.(events.EntityEvent))
	}
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
