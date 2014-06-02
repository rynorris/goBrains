package iomanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
)

// Manager is capable of handling multiple input/output
// paths at once. It abstracts this away from the main
// thread.
type Manager interface {
	// Shutdown turns of all IO.
	Shutdown()

	// Add registers an output channel with the Manager.
	Add(t IoType, out chan []DrawSpec)

	// Distribute sends the given entities out down all the output channels.
	Distribute(data []entity.Entity)

	// Handle is called by any input monitors to let the Manager deal with events.
	Handle(e events.Event)
}
