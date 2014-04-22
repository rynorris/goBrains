/*
 * Primitives are basic shapes which can be directly drawn to screen.
 *
 */

package graphics

import "image/color"
import "github.com/banthar/Go-SDL/sdl"

// A Primitive is a basic shape which can be drawn directly by the artist.
type Primitive interface {
	draw(s *sdl.Surface)
}

// A Point is as it sounds, a single point in space.
type Point struct {
	x, y int
	c    color.Color
}

// Points are drawn by setting a single corresponding pixel.
func (p Point) draw(s *sdl.Surface) {
	color := sdl.ColorFromGoColor(p.c)
	safeSet(s, p.x, p.y, color)
}

// A Rectangle is... a rectangle.
type Rectangle struct {
	x, y int16
	w, h uint16
	c    color.Color
}

// Rectangles are drawn by directly calling FillRect on the surface.
func (r Rectangle) draw(s *sdl.Surface) {
	format := s.Format
	color := sdl.ColorFromGoColor(r.c)
	colorVal := sdl.MapRGB(format, color.R, color.G, color.B)
	s.FillRect(&sdl.Rect{r.x, r.y, r.w, r.h}, colorVal)
}

// Circles are, you guessed it. Circles.
type Circle struct {
	x, y int16       // Location on screen
	r    uint16      // Radius
	b    int         // Border thickness. For now only controls if there IS a border or not, not actually it's thickness.
	c    color.Color // Color
}

// Circles may be filled or not.
func (c Circle) draw(s *sdl.Surface) {
	if c.b == 0 {
		drawFilledCircle(c.x, c.y, c.r, c.c, s)
	} else {
		drawOutlineCircle(c.x, c.y, c.r, c.c, s)
	}
}

// drawFilledCircle uses the integer midpoint circle algorithm to draw a filled
// circle to the given surface.
func drawFilledCircle(x0, y0 int16, r uint16, c color.Color, s *sdl.Surface) {
	format := s.Format
	color := sdl.ColorFromGoColor(c)
	colorVal := sdl.MapRGB(format, color.R, color.G, color.B)

	x := int16(r)
	y := int16(0)
	e := 1 - x

	for x >= y {
		s.FillRect(&sdl.Rect{-x + x0, y + y0, uint16(2 * x), 1}, colorVal)
		s.FillRect(&sdl.Rect{-x + x0, -y + y0, uint16(2 * x), 1}, colorVal)
		s.FillRect(&sdl.Rect{-y + x0, x + y0, uint16(2 * y), 1}, colorVal)
		s.FillRect(&sdl.Rect{-y + x0, -x + y0, uint16(2 * y), 1}, colorVal)

		y++

		if e < 0 {
			e += 2*y + 1
		} else {
			x--
			e += 2 * (y - x + 1)
		}
	}
}

// drawOutlineCircle uses the integer midpoint circle algorithm to draw the outline
// of a circle (1 px thick) to the given surface.
func drawOutlineCircle(x0, y0 int16, r uint16, c color.Color, s *sdl.Surface) {
	s.Lock()
	defer s.Unlock()

	color := sdl.ColorFromGoColor(c)

	x := int16(r)
	y := int16(0)
	e := 1 - x

	for x >= y {
		safeSet(s, int(x+x0), int(y+y0), color)
		safeSet(s, int(x+x0), int(-y+y0), color)
		safeSet(s, int(-x+x0), int(y+y0), color)
		safeSet(s, int(-x+x0), int(-y+y0), color)
		safeSet(s, int(y+x0), int(x+y0), color)
		safeSet(s, int(y+x0), int(-x+y0), color)
		safeSet(s, int(-y+x0), int(x+y0), color)
		safeSet(s, int(-y+x0), int(-x+y0), color)

		y++

		if e < 0 {
			e += 2*y + 1
		} else {
			x--
			e += 2 * (y - x + 1)
		}
	}
}

func safeSet(s *sdl.Surface, x, y int, c sdl.Color) {
	if x >= 0 && y >= 0 && x < int(s.W) && y < int(s.H) {
		s.Set(x, y, c)
	}
}
