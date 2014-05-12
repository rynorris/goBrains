// Definitions for different kinds of events.
// Note that Events in this package are GAME events, not
// SDL events.
// The job of this package is to take input from SDL, and
// convert it into game events which can be passed on into the
// main game itself.
package events

// Event interface.
// This is a little weird, but since there's no subclassing,
// we need an interface here to allow us to pass around
// different kinds of events blindly.
type Event interface {
	GetType() EventType // Returns the event type.
}

// A basic event.
// Contains just an event code, and no additional data.
type BasicEvent struct {
	EventType EventType
}

// GetType returns the event type code for the given event.
func (e BasicEvent) GetType() EventType {
	return e.EventType
}

// An event with an associated location.
type LocationEvent struct {
	BasicEvent
	x, y int
}
