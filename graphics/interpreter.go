/*
 * The Interpreter takes entities and breaks them down into component shapes
 * to be drawn by the lower-level graphics components.
 *
 */

package graphics

import "image/color"
import "github.com/discoviking/goBrains/entity"
import "github.com/discoviking/goBrains/food"

func Interpret(in chan entity.Entity, out chan Primitive) {
	for e := range in {
		switch e := e.(type) {
		case *food.Food:
			breakFood(e, out)
		default:
			breakEntity(e, out)
		}
	}
}

func breakEntity(e entity.Entity, out chan Primitive) {
	out <- Circle{20, 20, uint16(e.GetRadius()), 0, color.Black}
}

func breakFood(f *food.Food, out chan Primitive) {
	out <- Circle{20, 20, uint16(f.GetRadius()), 0, color.RGBA{50, 200, 50, 255}}
	out <- Circle{20, 20, uint16(f.GetRadius()), 1, color.Black}
}
