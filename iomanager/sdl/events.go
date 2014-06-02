package sdl

import (
	"github.com/DiscoViking/goBrains/events"
	"github.com/banthar/Go-SDL/sdl"
)

var (
	// Key Bindings
	keyDownBindings = map[uint32]events.EventType{
		sdl.K_z:      events.TOGGLE_DRAW,
		sdl.K_ESCAPE: events.TERMINATE,
		sdl.K_x:      events.TOGGLE_FRAME_LIMIT,
		sdl.K_MINUS:  events.SPEED_DOWN,
		sdl.K_EQUALS: events.SPEED_UP,
	}
)

// Converts SDL Events into game events.
func convert(e sdl.Event) events.Event {
	switch e := e.(type) {
	case *sdl.KeyboardEvent:
		// Keyboard event handling.
		switch e.Type {
		case sdl.KEYDOWN:
			// Key Down handling
			if t, ok := keyDownBindings[e.Keysym.Sym]; ok {
				return events.BasicEvent{t}
			}
		}
	}
	return nil
}
