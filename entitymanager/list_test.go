package entitymanager

import "testing"

type testEntity struct {
	dead bool
}

func (e *testEntity) Check() bool {
	return e.dead
}

func (e *testEntity) GetRadius() float64 { return 0 }
func (e *testEntity) Consume() float64   { return 0 }

func TestNewList(t *testing.T) {
	l := NewList()
	if len(l) != 0 {
		t.Errorf("Expected empty list, got %v entities.", len(l))
	}
}

func TestAdd(t *testing.T) {
	l := NewList()

	e := &testEntity{false}

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
		l.Add(&testEntity{false})
	}
	if len(l) != 6 {
		t.Errorf("Expected 6 entity in list, got %v.", len(l))
	}
}

func TestCheck(t *testing.T) {
	l := NewList()

	alive := &testEntity{false}
	dead := &testEntity{true}

	l.Add(alive)
	l.Add(dead)

	l.Check()

	if _, in := l[alive]; !in {
		t.Error("Live entity got removed from list.")
	}

	if _, in := l[dead]; in {
		t.Error("Dead entity didn't get removed from list.")
	}
}
