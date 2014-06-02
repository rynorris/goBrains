package sdl

import (
	"github.com/DiscoViking/goBrains/events"
	"github.com/banthar/Go-SDL/sdl"
)
import "testing"

// Test conversion of SDL keyboard events into game events
func TestConvertKeyboard(t *testing.T) {
	// Set up some event mappings we expect.
	tests := map[sdl.Event]events.Event{
		&sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_ESCAPE, 0, 0}}: events.BasicEvent{events.TERMINATE},

		&sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_z, 0, 0}}: events.BasicEvent{events.TOGGLE_DRAW},

		&sdl.KeyboardEvent{sdl.KEYUP, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_ESCAPE, 0, 0}}: nil,

		&sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_s, 0, 0}}: nil,
	}

	for in, out := range tests {
		if got := convert(in); got != out {
			t.Errorf("Input: %v, Expected: %v, Got: %v", in, out, got)
		}
	}
}
