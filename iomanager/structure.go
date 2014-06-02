package iomanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/locationmanager"
)

// DrawSpec is a combination of entity and location to send
// to anyone who might want to output it.
type DrawSpec struct {
	E   entity.Entity
	Loc locationmanager.Combination
}

type ioManager struct {
	lm     locationmanager.Location
	Out    map[IoType]chan []DrawSpec
	Events chan events.Event
}
