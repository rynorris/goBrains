package graphics

import (
	"image/color"
	"testing"

	"github.com/banthar/Go-SDL/sdl"
)

func TestArtist(t *testing.T) {
	draw := make(chan Primitive)
	done := make(chan struct{})
	defer close(done)

	s := sdl.CreateRGBSurface(0, 200, 200, 32, 0, 0, 0, 0)

	go Draw(draw, s, done)

	draw <- Circle{100, 100, 20, 0, color.RGBA{100, 100, 100, 255}}
	close(draw)
	<-done

	s.SaveBMP("test_output/TestArtist_got.bmp")
	CompareOutput("TestArtist", t)
}
