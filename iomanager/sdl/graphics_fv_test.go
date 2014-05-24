package sdl

import (
	"os"
	"testing"
	"time"

	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/iomanager"
	"github.com/DiscoViking/goBrains/locationmanager"
	"github.com/DiscoViking/goBrains/testutils"
	"github.com/banthar/Go-SDL/sdl"
)

func TestGraphicsFV(t *testing.T) {
	// This test does not run in Travis.
	if os.Getenv("TRAVIS") == "true" {
		t.Log("This test does not work in the Travis VMs. Passing by default.")
		return
	}

	lm := locationmanager.New()

	// Create some entities
	entities := []iomanager.DrawSpec{
		{food.New(lm, 1000), locationmanager.Combination{40, 40, 0}},
	}

	io := iomanager.New(lm)

	// We call mainLoop directly, since we need to be able
	// to access SDL later, and it can only run in one thread.
	data := make(chan []iomanager.DrawSpec, 1)

	// Send in the entities, close the channel.
	// Then when we call mainloop we should exit after one run.
	data <- entities

	go mainLoop(data, io)

	// Wait for drawing to finish
	<-time.After(1 * time.Second)
	// Get the screen, save it to an image and compare.
	s := sdl.GetVideoSurface()
	s.SaveBMP("test_output/TestGraphicsFV_got.bmp")

	testutils.CompareOutputImages("TestGraphicsFV", t)
}
