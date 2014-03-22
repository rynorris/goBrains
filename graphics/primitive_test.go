package graphics

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/discoviking/goBrains/testutils"
	"image/color"
	"os"
	"strconv"
	"testing"
)

func TestPoint(t *testing.T) {
	testname := "TestPoint"

	// Initialise SDL
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	s := sdl.CreateRGBSurface(0, 100, 100, 16, 0, 0, 0, 0)

	p := Point{15, 27, color.RGBA{28, 12, 231, 255}}
	p.draw(s)

	s.SaveBMP("test_output/" + testname + "_got.bmp")

	compareOutput(testname, t)
}

func TestCircle(t *testing.T) {
	// Initialise SDL
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	// Set up some circles to test.
	circles := []Circle{
		Circle{50, 50, 30, 0, color.RGBA{28, 12, 231, 255}},
		Circle{20, 10, 30, 0, color.RGBA{28, 12, 231, 255}},
		Circle{50, 50, 30, 1, color.RGBA{28, 12, 231, 255}},
		Circle{80, 70, 30, 1, color.RGBA{28, 12, 231, 255}},
	}

	ii := 1
	for _, c := range circles {
		testname := "TestCircle_" + strconv.Itoa(ii)

		s := sdl.CreateRGBSurface(0, 100, 100, 16, 0, 0, 0, 0)

		c.draw(s)

		s.SaveBMP("test_output/" + testname + "_got.bmp")

		compareOutput(testname, t)

		ii++
	}
}

func compareOutput(testname string, t *testing.T) {
	match, err := testutils.FilesAreEqual(
		"test_output/"+testname+"_got.bmp",
		"test_output/"+testname+"_exp.bmp")
	if err != nil {
		t.Errorf(err.Error())
	} else if !match {
		t.Errorf("Expected and actual outputs differ. Check files manually.")
	} else {
		//Pass, so remove _got file so we dont clog the output directory.
		os.Remove("test_output/" + testname + "_got.bmp")
	}
}
