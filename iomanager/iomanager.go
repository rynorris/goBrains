package iomanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	lm "github.com/DiscoViking/goBrains/locationmanager"
)

func New(lm lm.Location) *ioManager {
	out := make(map[IoType]chan []DrawSpec)
	event := make(chan events.Event, 5)
	io := &ioManager{lm, out, event}

	go func() {
		for e := range io.Events {
			events.Global.Broadcast(e)
		}
	}()

	return io
}

func (io *ioManager) Handle(e events.Event) {
	io.Events <- e
}

func (io *ioManager) Distribute(entities []entity.Entity) {
	specs := make([]DrawSpec, 0, len(entities))
	for _, e := range entities {
		spec, ok := Locate(io.lm, e)
		if !ok {
			continue
		}
		specs = append(specs, spec)
	}

	for _, out := range io.Out {
		select {
		case out <- specs:
		default:
		}
	}
}

func (io *ioManager) Add(t IoType, out chan []DrawSpec) {
	io.Stop(t)
	io.Out[t] = out
}

func (io *ioManager) Stop(t IoType) {
	out, ok := io.Out[t]
	if !ok {
		return
	}

	delete(io.Out, t)
	close(out)
}

func (io *ioManager) Shutdown() {
	for _, out := range io.Out {
		close(out)
	}
	close(io.Events)
}
