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

type IoManager struct {
	In     chan []entity.Entity
	Out    map[IoType]chan []DrawSpec
	Events chan events.Event
	Done   chan struct{}
}
