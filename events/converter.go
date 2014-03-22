package events

import "github.com/banthar/Go-SDL/sdl"

var (
	// Key Bindings
	terminate   = sdl.K_ESCAPE
	toggle_draw = sdl.KEY_z
)

// Listens for events from SDL, converts them into
// game events, and pipes them down the given channel
func Poll(out chan Event) {
	for {
		out <- convert(sdl.WaitEvent())
	}
}

// Converts SDL Events into game events.
func convert(e sdl.Event) Event {
	switch e.(type) {
	case sdl.KeyboardEvent:
		// Keyboard event handling.
		e := sdl.KeyboardEvent(e)
		switch e.Type {
		case sdl.KEYDOWN:
			// Key Down handling
			switch e.Keysym {
			case terminate:
				// Close the program
				out <- BasicEvent{TERMINATE}
			case toggle_draw:
				// Turn on/off drawing graphics
				out <- BasicEvent{TOGGLE_DRAW}
			}
		}
	}
}
