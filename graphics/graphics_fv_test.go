package graphics

import "github.com/DiscoViking/goBrains/entity"
import "github.com/DiscoViking/goBrains/food"
import "github.com/banthar/Go-SDL/sdl"
import "testing"

func TestGraphicsFV(t *testing.T) {
	data := make(chan []entity.Entity)
	done := make(chan struct{})

	// Create some entities
	entities := []entity.Entity{
		food.NewFood(1000),
	}

	go Start(data, done)

	// Get graphicsManager to draw them to the screen
	data <- entities
	<-done

	// Get the screen, save it to an image and compare.
	s := sdl.GetVideoSurface()
	s.SaveBMP("test_output/TestGraphicsFV_got.bmp")

	compareOutput("TestGraphicsFV", t)
}
