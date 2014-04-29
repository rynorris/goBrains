package events

import "github.com/banthar/Go-SDL/sdl"

var (
	// Key Bindings
	keyDownBindings = map[uint32]EventType{
		sdl.K_z:      TOGGLE_DRAW,
		sdl.K_ESCAPE: TERMINATE,
	}
)

// Converts SDL Events into game events.
func Convert(e sdl.Event) Event {
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
