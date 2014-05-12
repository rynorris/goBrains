package events

import "testing"

func TestRegister(t *testing.T) {
	h := NewHandler()

	l := func(Event) { return }
	h.Register(TERMINATE, l)

	if len(h[TERMINATE]) != 1 {
		t.Errorf("Registered listener not in listeners map.")
	}
}

func TestBroadcast(t *testing.T) {
	h := NewHandler()

	var c = 0
	l := func(Event) { c += 1 }
	h.Register(TERMINATE, l)

	h.Broadcast(BasicEvent{TERMINATE})

	if c != 1 {
		t.Fatal("Listening fuction was not called when it's event type was broadcast.")
	}

	h.Broadcast(BasicEvent{NONE})

	if c != 1 {
		t.Fatal("Listening function was called when a different event type was broadcase.")
	}
}
