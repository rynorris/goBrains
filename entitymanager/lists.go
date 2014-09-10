package entitymanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
)

type entityList map[entity.Entity]struct{}

func (l entityList) Work() {
	workers := 16
	work := make(chan entity.Entity)
	done := make(chan struct{})
	defer close(done)

	for i := 0; i < workers; i++ {
		go func() {
			for e := range work {
				e.Work()
			}
			done <- struct{}{}
		}()
	}

	for e, _ := range l {
		work <- e
	}

	close(work)

	for i := 0; i < workers; i++ {
		<-done
	}
}
func (l entityList) Check() {
	for e, _ := range l {
		if e.Check() {
			// Raise an entity destroy event.
			events.Global.Broadcast(events.EntityEvent{
				Type: events.ENTITY_DESTROY,
				E:    e,
			})
			delete(l, e)
		}
	}
}

func (l entityList) Clear() {
	for e, _ := range l {
		delete(l, e)
	}
}

func (l entityList) Add(e entity.Entity) {
	// Raise an entity create event.
	events.Global.Broadcast(events.EntityEvent{
		Type: events.ENTITY_CREATE,
		E:    e,
	})
	l[e] = struct{}{}
}

func (l entityList) Slice() []entity.Entity {
	s := make([]entity.Entity, 0, len(l))
	for e, _ := range l {
		s = append(s, e)
	}
	return s
}

func NewList() entityList {
	return map[entity.Entity]struct{}{}
}
