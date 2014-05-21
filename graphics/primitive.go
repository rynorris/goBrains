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

// Specifies a line to be drawn.
type Line struct {
	x0, y0, x1, y1 int16
	c              color.Color
}

func (l Line) draw(s *sdl.Surface) {
	drawLine(l.x0, l.y0, l.x1, l.y1, l.c, s)
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

// Uses Bresenham's algorithm to draw a line between two points.
func drawLine(x0, y0, x1, y1 int16, c color.Color, s *sdl.Surface) {
	s.Lock()
	defer s.Unlock()

	color := sdl.ColorFromGoColor(c)

	// Make sure the two ends are left-to-right.
	if x1 < x0 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	// This algorithm only works for curves where dx > -dy and dy < 0
	// So, prepare a coordinate transform to make this the case.
	// We will then reverse the transform when we plot the points.
	// There are 4 cases, all the transformations are self-inverse, which
	// makes our lives a little easier.
	dx := int(x1 - x0)
	dy := int(y1 - y0)
	var transform func(x, y int) (int, int)
	var inverse func(x, y int) (int, int)

	if dy < 0 {
		if dx < -dy {
			transform = func(x, y int) (int, int) { return -y, x }
			inverse = func(x, y int) (int, int) { return y, -x }
		} else {
			transform = func(x, y int) (int, int) { return x, -y }
			inverse = transform
		}
	} else {
		if dx < dy {
			transform = func(x, y int) (int, int) { return y, x }
			inverse = transform
		} else {
			transform = func(x, y int) (int, int) { return x, y }
			inverse = transform
		}
	}

	// Transform coordinates.
	tx0, ty0 := transform(int(x0), int(y0))
	tx1, ty1 := transform(int(x1), int(y1))

	// Recalculate dx and dy.
	dx = tx1 - tx0
	dy = ty1 - ty0

	D := 2*dy - dx

	safeSet(s, int(x0), int(y0), color)

	y := ty0

	for x := tx0 + 1; x <= tx1; x++ {
		if D > 0 {
			y += 1
			D += 2*dy - 2*dx
		} else {
			D += 2 * dy
		}
		tx, ty := inverse(x, y)
		safeSet(s, tx, ty, color)
	}
}

func safeSet(s *sdl.Surface, x, y int, c sdl.Color) {
	if x >= 0 && y >= 0 && x < int(s.W) && y < int(s.H) {
		s.Set(x, y, c)
	}
}
