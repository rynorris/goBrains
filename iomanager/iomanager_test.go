package iomanager

import (
	"testing"

	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/locationmanager"
)

func TestHandle(t *testing.T) {
	lm := locationmanager.New()
	io := New(lm)
	defer io.Shutdown()

	received := false
	wait := make(chan struct{})
	events.Global.Register(events.TERMINATE, func(events.Event) {
		received = true
		wait <- struct{}{}
	})

	ev := events.BasicEvent{events.TERMINATE}

	io.Handle(ev)
	<-wait

	if !received {
		t.Errorf("Didn't receive event.")
	}
}

func TestDistribute(t *testing.T) {
	lm := locationmanager.New()
	io := New(lm)

	out := make(chan []DrawSpec, 1)

	io.Add(SDL, out)

	e1 := &entity.TestEntity{}
	e2 := &entity.TestEntity{}

	// Only e1 is locatable
	lm.AddEntAtLocation(e1, locationmanager.Combination{0, 0, 0})
	entities := []entity.Entity{e1, e2}

	io.Distribute(entities)

	exp := []DrawSpec{{e1, locationmanager.Combination{0, 0, 0}}}
	got := <-out

	if len(exp) != len(got) {
		t.Fatalf("Expected and received DrawSpec slices of different length. Exp: %v, Got: %v", len(exp), len(got))
	}

	for i := 0; i < len(exp); i++ {
		if exp[i] != got[i] {
			t.Errorf("Wrong DrawSpec at index %v. Exp: %v, Got: %v", i, exp[i], got[i])
		}
	}
}

func TestStop(t *testing.T) {
	out := make(chan []DrawSpec)
	done := make(chan struct{})

	lm := locationmanager.New()
	io := New(lm)

	// Send off a goroutine waiting on this channel.
	go func() {
		<-out
		done <- struct{}{}
	}()

	io.Add(SDL, out)
	// Stop it. This should cause the earlier goroutine to unblock.
	io.Stop(SDL)

	// Receive the signal that we've unblocked.
	// If we hang here, all is lost.
	<-done
}

func TestShutdown(t *testing.T) {
	out := make(chan []DrawSpec)
	done := make(chan struct{})

	lm := locationmanager.New()
	io := New(lm)

	// Send off a goroutine waiting on the two channels which should get closed.
	go func() {
		<-out
		done <- struct{}{}
		<-io.Events
		done <- struct{}{}
	}()

	io.Add(SDL, out)

	io.Shutdown()

	// Listen twice on done. Our goroutine should clear out and events.
	<-done
	<-done
}
