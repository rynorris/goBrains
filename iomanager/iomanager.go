package iomanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	lm "github.com/DiscoViking/goBrains/locationmanager"
)

func New(lm lm.Location) *IoManager {
	data := make(chan []entity.Entity)
	out := make(map[IoType]chan []DrawSpec)
	event := make(chan events.Event, 5)
	done := make(chan struct{})
	io := &IoManager{data, out, event, done}
	go io.Distribute(lm, done)

	go func() {
		for e := range io.Events {
			events.Global.Broadcast(e)
		}
	}()

	return io
}

func (io *IoManager) Distribute(
	lm lm.Location,
	done chan struct{}) {

	for entities := range io.In {
		specs := make([]DrawSpec, 0, len(entities))
		for _, e := range entities {
			spec, ok := Locate(lm, e)
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
		done <- struct{}{}
	}
}

func (io *IoManager) Stop(t IoType) {
	out, ok := io.Out[t]
	if !ok {
		return
	}

	delete(io.Out, t)
	close(out)
}

func (io *IoManager) Shutdown() {
	for _, out := range io.Out {
		close(out)
	}
}
