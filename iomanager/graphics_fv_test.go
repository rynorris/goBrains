package iomanager

import (
	"fmt"
	"os"
	"testing"

	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/food"
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

	data := make(chan []entity.Entity)
	done := make(chan struct{})
	event := make(chan sdl.Event)

	events.Global.Register(events.TERMINATE,
		func(events.Event) { close(data) })

	go func() {
		for e := range event {
			fmt.Println("Got event!")
			events.Handle(e)
		}
	}()

	lm := locationmanager.New()

	// Create some entities
	entities := []entity.Entity{
		food.New(lm, 1000),
	}

	go Start(lm, data, done, event)

	// Get graphicsManager to draw them to the screen
	data <- entities
	<-done

	// Get the screen, save it to an image and compare.
	s := sdl.GetVideoSurface()
	s.SaveBMP("test_output/TestGraphicsFV_got.bmp")

	testutils.CompareOutputImages("TestGraphicsFV", t)
}
