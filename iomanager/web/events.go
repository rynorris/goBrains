package web

import "github.com/DiscoViking/goBrains/events"

type event struct {
	Type        string
	Key         string
	MouseButton int
	X           int
	Y           int
}

var keyBindings = map[string]events.EventType{
	"Z": events.TOGGLE_FRAME_LIMIT,
	"Q": events.SPEED_DOWN,
	"W": events.SPEED_UP,
}

func convert(e event) events.Event {
	switch e.Type {
	case "Key":
		if t, ok := keyBindings[e.Key]; ok {
			return events.BasicEvent{t}
		}
	}

	return nil
}
