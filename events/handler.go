package events

// listener is a function to be called whenever
// an event is raised.
type listener func(Event)

// Handler farms events out to anyone who has registed interest.
type Handler map[EventType][]listener

// NewHandler returns a new Handler with no registered listeners.
func NewHandler() Handler {
	return Handler(map[EventType][]listener{})
}

// Register ensures the listener will get called whenever
// an event of type t is raised.
func (h Handler) Register(t EventType, l listener) {
	h[t] = append(h[t], l)
}

// Broadcast sends an event out to any listeners who care.
func (h Handler) Broadcast(e Event) {
	for _, l := range h[e.GetType()] {
		l(e)
	}
}

// Clear all listeners.
func (h Handler) Reset() {
	for k, _ := range h {
		h[k] = []listener{}
	}
}
