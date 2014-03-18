/*
 * The Interpreter takes entities and breaks them down into component shapes
 * to be drawn by the lower-level graphics components.
 *
 */

package graphics

import "github.com/discoviking/goBrains/entity"

func Interpret(in chan entity.Entity, out chan Primitive) {
	for e := range in {
		switch e.type {
		default:
			out <- breakEntity(e)
		}
	}
}

func breakEntity(e entity.Entity) Primitive {
	return Point{10, 10, color.Black}
}
