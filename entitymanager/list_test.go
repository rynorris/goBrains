package entitymanager

import (
	"reflect"
	"testing"

	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
)

func TestNewList(t *testing.T) {
	l := NewList()
	if len(l) != 0 {
		t.Errorf("Expected empty list, got %v entities.", len(l))
	}
}

func TestAdd(t *testing.T) {
	l := NewList()

	e := &entity.TestEntity{}

	// Create function to report if entity event contains the wrong entity.
	defer events.Global.Reset()
	events.Global.Register(events.ENTITY_CREATE, func(ev events.Event) {
		eev := ev.(events.EntityEvent)
		if !reflect.DeepEqual(eev.E, e) {
			t.Errorf("Received ENTITY_CREATE event for wrong entity. Expected: %v, Got: %v.\n", e, eev.E)
		}
	})
	events.Global.Register(events.ENTITY_DESTROY, func(ev events.Event) {
		t.Errorf("Received ENTITY_DESTROY.")
	})

	t.Log("Adding entity to list.")
	l.Add(e)
	if len(l) != 1 {
		t.Errorf("Expected 1 entity in list, got %v.", len(l))
	}

	t.Log("Adding duplicate entity to list. Should not get added twice.")
	l.Add(e)
	if len(l) != 1 {
		t.Errorf("Expected 1 entity in list, got %v.", len(l))
	}

	t.Log("Adding 5 more entities to list.")
	for i := 0; i < 5; i++ {
		l.Add(&entity.TestEntity{})
	}
	if len(l) != 6 {
		t.Errorf("Expected 6 entity in list, got %v.", len(l))
	}
}

func TestClear(t *testing.T) {
	l := NewList()

	l.Add(&entity.TestEntity{})
	l.Add(&entity.TestEntity{})

	if len(l) != 2 {
		t.Errorf("Added 2 entities to list, but length is %v.", len(l))
	}

	l.Clear()

	if len(l) != 0 {
		t.Errorf("After clearing list, still contained %v entities.\nList contents: %#v", len(l), l)
	}
}

func TestCheck(t *testing.T) {
	l := NewList()

	alive := &entity.TestEntity{}
	dead := &entity.TestEntity{}
	dead.TeDead = true

	l.Add(alive)
	l.Add(dead)

	// Create function to report if entity event contains the wrong entity.
	defer events.Global.Reset()
	events.Global.Register(events.ENTITY_DESTROY, func(ev events.Event) {
		eev := ev.(events.EntityEvent)
		if !reflect.DeepEqual(eev.E, dead) {
			t.Errorf("Received ENTITY_DESTROY for %v instead of %v.", eev.E, dead)
		}
	})

	l.Check()

	if _, in := l[alive]; !in {
		t.Error("Live entity got removed from list.")
	}

	if _, in := l[dead]; in {
		t.Error("Dead entity didn't get removed from list.")
	}
}
