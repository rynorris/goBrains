package graphics

import (
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/locationmanager"
	"image/color"
	"reflect"
	"testing"
	"time"
)

// Test that the interpreter does what we expect when it's given an
// entity it doesn't recognise.
func TestInterpretDefault(t *testing.T) {
	in := make(chan entity.Entity)
	out := make(chan Primitive)
	defer close(in)

	lm := locationmanager.New()

	go Interpret(lm, in, out)

	e := &entity.TestEntity{}
	e.TeRadius = 10
	lm.AddEntity(e)

	_, loc := lm.GetLocation(e)
	expected := Circle{int16(loc.X), int16(loc.Y), 10, 0, color.RGBA{255, 255, 255, 255}}

	in <- e

	output := <-out

	// Test it output a circle.
	switch T := output.(type) {
	case Circle:
		// Do Nothing, this is correct
	default:
		t.Errorf("Expected circle, got %v.", T)
	}

	circle := output.(Circle)

	// Test the circle was what we expected.
	if circle != expected {
		t.Errorf("Expected circle x=%v y=%v r=%v c=%v\n"+
			"Got x=%v y=%v r=%v c=%v",
			expected.x, expected.y, expected.r, expected.c,
			circle.x, circle.y, circle.r, circle.c)
	}
}

// Test the interpreter does what we expect when given food.
func TestInterpretFood(t *testing.T) {
	in := make(chan entity.Entity)
	out := make(chan Primitive)
	defer close(in)

	lm := locationmanager.New()

	go Interpret(lm, in, out)

	f := food.New(lm, 100)

	in <- f

	_, loc := lm.GetLocation(f)
	expected := Circle{int16(loc.X), int16(loc.Y), 10, 0, color.RGBA{50, 200, 50, 255}}

	output := <-out

	// Test it output a Circle.
	switch T := output.(type) {
	case Circle:
		// Do Nothing, this is correct
	default:
		t.Errorf("Expected circle, got %v.", T)
	}

	circle := output.(Circle)

	// Test the circle was what we expected.
	if circle != expected {
		t.Errorf("Expected circle x=%v y=%v r=%v c=%v\n"+
			"Got x=%v y=%v r=%v c=%v",
			expected.x, expected.y, expected.r, expected.c,
			circle.x, circle.y, circle.r, circle.c)
	}
}

// Test the interpreter does what we expect when given a creature.
func TestInterpretCreature(t *testing.T) {
	in := make(chan entity.Entity)
	out := make(chan Primitive)
	defer close(in)

	lm := locationmanager.New()
	lm.StartAtOrigin()

	go Interpret(lm, in, out)

	c := creature.NewSimple(lm)

	in <- c

	expected := []Primitive{
		Line{0, 0, 34, 20, color.RGBA{170, 170, 170, 255}},
		Line{0, 0, 34, -20, color.RGBA{170, 170, 170, 255}},
		Circle{0, 0, 8, 0, color.RGBA{200, 50, 50, 255}},
		Circle{-6, 0, 6, 0, color.RGBA{200, 50, 50, 255}},
		Circle{-10, 0, 4, 0, color.RGBA{200, 50, 50, 255}},
		Circle{6, 0, 2, 0, color.Black},
		Circle{34, 20, 2, 0, color.RGBA{200, 200, 50, 255}},
		Circle{34, -20, 2, 0, color.RGBA{200, 200, 50, 255}},
	}

	timeout := time.After(5 * time.Second)

	for _, p := range expected {
		select {
		case got := <-out:
			if !reflect.DeepEqual(p, got) {
				t.Errorf("Expected %v, got %v\n", p, got)
			}
		case <-timeout:
			t.Errorf("Timed out. Not enough primitives output.\n")
		}
	}

	// Wait a second to see if the interpreter spits out any
	// erroneous extra output.
	timeout = time.After(1 * time.Second)

loop:
	for {
		select {
		case extra := <-out:
			t.Errorf("Got extra output: %#v\n", extra)
		case <-timeout:
			break loop
		}
	}
}
