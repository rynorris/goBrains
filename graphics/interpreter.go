/*
 * The Interpreter takes entities and breaks them down into component shapes
 * to be drawn by the lower-level graphics components.
 *
 */

package graphics

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/locationmanager"
)

func Interpret(lm locationmanager.Location, in chan entity.Entity, out chan Primitive) {
	defer close(out)
	for e := range in {
		switch e := e.(type) {
		case *food.Food:
			breakFood(lm, e, out)
		default:
			breakEntity(lm, e, out)
		}
	}
}

func breakEntity(lm locationmanager.Location, e entity.Entity, out chan Primitive) {
	ok, comb := lm.GetLocation(e)
	if !ok {
		return
	}
	x, y := comb.X, comb.Y

	out <- Circle{int16(x), int16(y), uint16(10), 0, e.GetColor()}
}

func breakFood(lm locationmanager.Location, f *food.Food, out chan Primitive) {
	ok, comb := lm.GetLocation(f)
	if !ok {
		return
	}
	x, y := comb.X, comb.Y
	out <- Circle{int16(x), int16(y), uint16(f.GetRadius()), 0, f.GetColor()}
}
