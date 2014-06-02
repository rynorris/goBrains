package entitymanager

import (
	"testing"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/entity"
)

func TestNewList(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	l := NewList()
	if len(l) != 0 {
		t.Errorf("Expected empty list, got %v entities.", len(l))
	}
}

func TestAdd(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	l := NewList()

	e := &entity.TestEntity{}

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
	config.Load("../config/test_config.gcfg")
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
	config.Load("../config/test_config.gcfg")
	l := NewList()

	alive := &entity.TestEntity{}
	dead := &entity.TestEntity{}
	dead.TeDead = true

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
