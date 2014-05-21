// Unit tests for the primitive drawing module.
// These work by drawing a primitve to an SDL surface,
// saving it to a .bmp, and then comparing that file
// with a pre-saved .bmp file in the test_output/ directory.
//
// If you change any of the tests here so that they expect different output,
// you will need to update the expected output images.
// This can be done by running the tests (so they will fail),
// any failed tests will leave a TestName_got.bmp in the
// test_output/ folder.
// CHECK THAT THIS OUTPUT IS CORRECT.
// Then rename the image to TestName_exp.bmp.
//
// The test should now pass.

package graphics

import (
	"image/color"
	"os"
	"strconv"
	"testing"

	"github.com/banthar/Go-SDL/sdl"
)

func TestPoint(t *testing.T) {
	// This test does not run in Travis.
	if os.Getenv("TRAVIS") == "true" {
		t.Log("This test does not work in the Travis VMs. Passing by default.")
		return
	}

	testname := "TestPoint"

	// Initialise SDL
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	s := sdl.CreateRGBSurface(0, 100, 100, 16, 0, 0, 0, 0)

	p := Point{15, 27, color.RGBA{28, 12, 231, 255}}
	p.draw(s)

	s.SaveBMP("test_output/" + testname + "_got.bmp")

	CompareOutput(testname, t)
}

func TestRectangle(t *testing.T) {
	// This test does not run in Travis.
	if os.Getenv("TRAVIS") == "true" {
		t.Log("This test does not work in the Travis VMs. Passing by default.")
		return
	}

	// Initialise SDL
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	// Set up some circles to test
	rects := []Rectangle{
		Rectangle{50, 50, 30, 20, color.RGBA{28, 12, 231, 255}}, // Basic filled rectangle
		Rectangle{50, 80, 20, 40, color.RGBA{28, 12, 231, 255}}, // Rectangle clipped by edge of screen
	}

	ii := 1
	for _, r := range rects {
		testname := "TestRectangle_" + strconv.Itoa(ii)

		s := sdl.CreateRGBSurface(0, 100, 100, 16, 0, 0, 0, 0)

		r.draw(s)

		s.SaveBMP("test_output/" + testname + "_got.bmp")

		CompareOutput(testname, t)

		ii++
	}
}

func TestCircle(t *testing.T) {
	// This test does not run in Travis.
	if os.Getenv("TRAVIS") == "true" {
		t.Log("This test does not work in the Travis VMs. Passing by default.")
		return
	}

	// Initialise SDL
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	// Set up some circles to test
	circles := []Circle{
		Circle{50, 50, 30, 0, color.RGBA{28, 12, 231, 255}}, // Basic filled circle
		Circle{20, 10, 30, 0, color.RGBA{28, 12, 231, 255}}, // Filled circle, clipped by edge of surface
		Circle{50, 50, 30, 1, color.RGBA{28, 12, 231, 255}}, // Basic outline circle
		Circle{80, 70, 30, 1, color.RGBA{28, 12, 231, 255}}, // Outling circle, clipper by edge of surface
	}

	ii := 1
	for _, c := range circles {
		testname := "TestCircle_" + strconv.Itoa(ii)

		s := sdl.CreateRGBSurface(0, 100, 100, 16, 0, 0, 0, 0)

		c.draw(s)

		s.SaveBMP("test_output/" + testname + "_got.bmp")

		CompareOutput(testname, t)

		ii++
	}
}

func TestLine(t *testing.T) {
	// This test does not run in Travis.
	if os.Getenv("TRAVIS") == "true" {
		t.Log("This test does not work in the Travis VMs. Passing by default.")
		return
	}

	// Initialise SDL
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	// Set up some lines to test
	lines := []Line{
		Line{10, 10, 20, 80, color.RGBA{28, 12, 231, 255}}, // Steep downwards.
		Line{10, 60, 70, 80, color.RGBA{28, 12, 231, 255}}, // Shallow downwards.
		Line{10, 30, 70, 20, color.RGBA{28, 12, 231, 255}}, // Shallow upwards.
		Line{10, 70, 30, 20, color.RGBA{28, 12, 231, 255}}, // Steep upwards.
		Line{10, 30, 10, 80, color.RGBA{28, 12, 231, 255}}, // Vertical.
		Line{10, 20, 60, 20, color.RGBA{28, 12, 231, 255}}, // Horizontal.
	}

	for i, l := range lines {
		testname := "TestLine_" + strconv.Itoa(i)

		s := sdl.CreateRGBSurface(0, 100, 100, 16, 0, 0, 0, 0)

		l.draw(s)

		s.SaveBMP("test_output/" + testname + "_got.bmp")

		CompareOutput(testname, t)
	}
}
