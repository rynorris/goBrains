package events

import "github.com/banthar/Go-SDL/sdl"
import "testing"

// Test conversion of SDL keyboard events into game events
func TestConvertKeyboard(t *testing.T) {
	// Set up some event mappings we expect.
	tests := map[sdl.Event]Event{
		&sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_ESCAPE, 0, 0}}: BasicEvent{TERMINATE},

		&sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0},
			sdl.Keysym{0, sdl.K_z, 0, 0}}: BasicEvent{TOGGLE_DRAW},

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

func TestHandle(t *testing.T) {
	c := 0
	Global.Register(TOGGLE_DRAW, func(e Event) { c += 1 })

	Handle(&sdl.KeyboardEvent{sdl.KEYDOWN, 0, 0, [1]byte{0}, sdl.Keysym{0, sdl.K_z, 0, 0}})

	if c != 1 {
		t.Error("Callback function not called on event.")
	}
}
