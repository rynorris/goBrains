package events

import "github.com/banthar/Go-SDL/sdl"

var (
	// Key Bindings
	keyDownBindings = map[uint32]EventType{
		sdl.K_z:      TOGGLE_DRAW,
		sdl.K_ESCAPE: TERMINATE,
	}
)

// Listens for events from SDL, converts them into
// game events, and pipes them down the given channel
func Poll(out chan Event) {
	for {
		if e := convert(sdl.WaitEvent()); e.GetType() != NONE {
			out <- e
		}
	}
}

// Converts SDL Events into game events.
func convert(e sdl.Event) Event {
	switch e := e.(type) {
	case sdl.KeyboardEvent:
		// Keyboard event handling.
		switch e.Type {
		case sdl.KEYDOWN:
			// Key Down handling
			if t, ok := keyDownBindings[e.Keysym.Sym]; ok {
				return BasicEvent{t}
			}
		}
	}
	return BasicEvent{NONE}
}
