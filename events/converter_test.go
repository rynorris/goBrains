package events

import "github.com/banthar/Go-SDL/sdl"
import "testing"

// Test conversion of SDL keyboard events into game events
func TestConvertKeyboard(t *testing.T) {
	// Set up some event mappings we expect.
	tests := map[sdl.Event]Event{
		sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_ESCAPE, 0, 0}}: BasicEvent{TERMINATE},

		sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_z, 0, 0}}: BasicEvent{TOGGLE_DRAW},

		sdl.KeyboardEvent{sdl.KEYUP, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_ESCAPE, 0, 0}}: BasicEvent{NONE},

		sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_s, 0, 0}}: BasicEvent{NONE},
	}

	for in, out := range tests {
		if got := Convert(in); got != out {
			t.Errorf("Input: %v, Expected: %v, Got: %v", in, out, got)
		}
	}
}
