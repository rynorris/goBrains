package events

import "github.com/banthar/Go-SDL/sdl"

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
			case sdl.K_ESCAPE:
				// Escape key
				out <- BasicEvent{TERMINATE}
			case sdl.K_z:
				// Z key
				out <- BasicEvent{TOGGLE_DRAW}
			}
		}
	}
}
