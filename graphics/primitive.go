/*
 * Primitives are basic shapes which can be directly drawn to screen.
 *
 */

package graphics

import "math"
import "image/color"
import "github.com/banthar/Go-SDL/sdl"

type Primitive interface {
    draw(s *sdl.Surface)
}

type Point struct {
    x, y int
    c    color.Color
}

func (p Point) draw(s *sdl.Surface) {
    s.Set(p.x, p.y, p.c)
}

type Rectangle struct {
    x, y int16
    w, h uint16
    c    color.Color
}

func (r Rectangle) draw(s *sdl.Surface) {
    format := sdl.GetVideoInfo().Vfmt
    color := sdl.ColorFromGoColor(r.c)
    colorVal := sdl.MapRGB(format, color.R, color.G, color.B)
    s.FillRect(&sdl.Rect{r.x, r.y, r.w, r.h}, colorVal)
}

type Circle struct {
    x, y int16
    r    uint16
    c    color.Color
}

func (c Circle) draw(s *sdl.Surface) {
    format := sdl.GetVideoInfo().Vfmt
    color := sdl.ColorFromGoColor(c.c)
    colorVal := sdl.MapRGB(format, color.R, color.G, color.B)

    for h := 0; h <= int(c.r); h++ {
        l := int16(math.Sqrt(float64(int(c.r*c.r) - (h * h))))

        x := int16(c.x) - l

        s.FillRect(&sdl.Rect{x, c.y + int16(h), uint16(2 * l), 1}, colorVal)
        s.FillRect(&sdl.Rect{x, c.y - int16(h), uint16(2 * l), 1}, colorVal)
    }
}
