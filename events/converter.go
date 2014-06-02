package events

import "github.com/banthar/Go-SDL/sdl"

var (
	// Key Bindings
	keyDownBindings = map[uint32]EventType{
		sdl.K_z:      TOGGLE_DRAW,
		sdl.K_ESCAPE: TERMINATE,
		sdl.K_x:      TOGGLE_FRAME_LIMIT,
		sdl.K_EQUALS: SPEED_UP,
		sdl.K_MINUS:  SPEED_DOWN,
	}
)

// Converts SDL Events into game events.
func convert(e sdl.Event) Event {
	switch e := e.(type) {
	case *sdl.KeyboardEvent:
		// Keyboard event handling.
		switch e.Type {
		case sdl.KEYDOWN:
			// Key Down handling
			if t, ok := keyDownBindings[e.Keysym.Sym]; ok {
				return BasicEvent{t}
			}
		}
	}
	return nil
}

// Handle converts an SDL event to a game event, and
// broadcasts them to anyone listening on the global channel.
func Handle(e sdl.Event) {
	ev := convert(e)
	if ev != nil {
		Global.Broadcast(ev)
	}
}
