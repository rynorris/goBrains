/*
 * The Artist is the component who actually draws to the screen.
 *
 * It only understands Primitive shapes however, so must be first assisted by
 * an interpreter to break down more complex objects.
 */

package graphics

import "github.com/banthar/Go-SDL/sdl"

func Draw(in chan Primitive, s *sdl.Surface, done chan struct{}) {
	for p := range in {
		p.draw(s)
	}
	done <- struct{}{}
}
