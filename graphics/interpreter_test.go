package graphics

import "testing"
import "image/color"
import "github.com/DiscoViking/goBrains/entity"
import "github.com/DiscoViking/goBrains/food"
import "github.com/DiscoViking/goBrains/locationmanager"

// A dummy entity we use for testing the interpreter.
type testEntity struct {
}

// The testEntity always has radius 10.
func (t testEntity) GetRadius() float64 {
	return 10
}

// Test that the interpreter does what we expect when it's given an
// entity it doesn't recognise.
func TestInterpretDefault(t *testing.T) {
	in := make(chan entity.Entity)
	out := make(chan Primitive)
	defer close(in)

	go Interpret(in, out)

	expected := Circle{100, 100, 10, 0, color.Black}

	in <- &testEntity{}

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

	go Interpret(in, out)

	lm := locationmanager.NewLocationManager()

	in <- food.New(lm, 100)

	expected := Circle{100, 100, 11, 0, color.Black}

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

	expected = Circle{100, 100, 10, 0, color.RGBA{50, 200, 50, 255}}

	output = <-out

	// Test it output a Circle.
	switch T := output.(type) {
	case Circle:
		// Do Nothing, this is correct
	default:
		t.Errorf("Expected circle, got %v.", T)
	}

	circle = output.(Circle)

	// Test the circle was what we expected.
	if circle != expected {
		t.Errorf("Expected circle x=%v y=%v r=%v c=%v\n"+
			"Got x=%v y=%v r=%v c=%v",
			expected.x, expected.y, expected.r, expected.c,
			circle.x, circle.y, circle.r, circle.c)
	}
}
