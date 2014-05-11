/*
 * The Interpreter takes entities and breaks them down into component shapes
 * to be drawn by the lower-level graphics components.
 *
 */

package graphics

import "image/color"

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/locationmanager"
)
import "github.com/DiscoViking/goBrains/food"

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
	ok, x, y, _ := lm.GetLocation(e)
	if !ok {
		return
	}

	out <- Circle{int16(x), int16(y), uint16(10), 0, color.Black}
}

func breakFood(lm locationmanager.Location, f *food.Food, out chan Primitive) {
	ok, x, y, _ := lm.GetLocation(f)
	if !ok {
		return
	}
	out <- Circle{int16(x), int16(y), uint16(f.GetRadius()), 0, color.RGBA{50, 200, 50, 255}}
}
