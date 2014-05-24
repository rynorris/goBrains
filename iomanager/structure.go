package iomanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/locationmanager"
)

type DrawSpec struct {
	E   entity.Entity
	Loc locationmanager.Combination
}

type ioManager struct {
	lm     locationmanager.Location
	Out    map[IoType]chan []DrawSpec
	Events chan events.Event
	Done   chan struct{}
}

type Manager interface {
	Shutdown()
	Add(t IoType, out chan []DrawSpec)
	Distribute(data []entity.Entity)
	Handle(e events.Event)
}
