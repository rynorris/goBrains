/*
 * Primitives are basic shapes which can be directly drawn to screen.
 *
 */

package graphics

import "image/color"
import "github.com/banthar/Go-SDL/sdl"

type Primitive interface {
	draw(s *Surface)
}

type Point struct {
	x, y int
	c    color.Color
}

func (p Point) draw(s *sdl.Surface) {
	s.Set(x, y, c)
}
