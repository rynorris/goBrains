/*
 * The Interpreter takes entities and breaks them down into component shapes
 * to be drawn by the lower-level graphics components.
 *
 */

package graphics

import (
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/locationmanager"
	"image/color"
	"math"
)

func Interpret(lm locationmanager.Location, in chan entity.Entity, out chan Primitive) {
	defer close(out)
	for e := range in {
		switch e := e.(type) {
		case *creature.Creature:
			breakCreature(lm, e, out)
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

func breakCreature(lm locationmanager.Location, c *creature.Creature, out chan Primitive) {
	ok, comb := lm.GetLocation(c)
	if !ok {
		return
	}
	var dx, dy float64

	x, y, o := comb.X, comb.Y, comb.Orient
	cosO := math.Cos(o)
	sinO := math.Sin(o)

	// Draw the antenna lines first, so that the circles cover them.
	dx = math.Cos(o+math.Pi/6) * 40
	dy = math.Sin(o+math.Pi/6) * 40
	out <- Line{int16(x), int16(y), int16(x + dx), int16(y + dy), color.RGBA{170, 170, 170, 255}}
	dx = math.Cos(o-math.Pi/6) * 40
	dy = math.Sin(o-math.Pi/6) * 40
	out <- Line{int16(x), int16(y), int16(x + dx), int16(y + dy), color.RGBA{170, 170, 170, 255}}

	// Body
	out <- Circle{int16(x), int16(y), uint16(8), 0, color.RGBA{200, 50, 50, 255}}
	dx = cosO * 6
	dy = sinO * 6
	out <- Circle{int16(x - dx), int16(y - dy), uint16(6), 0, color.RGBA{200, 50, 50, 255}}
	dx = cosO * 10
	dy = sinO * 10
	out <- Circle{int16(x - dx), int16(y - dy), uint16(4), 0, color.RGBA{200, 50, 50, 255}}

	// Mouth
	dx = cosO * 6
	dy = sinO * 6
	out <- Circle{int16(x + dx), int16(y + dy), uint16(2), 0, color.Black}

	// Antennae
	dx = math.Cos(o+math.Pi/6) * 40
	dy = math.Sin(o+math.Pi/6) * 40
	out <- Circle{int16(x + dx), int16(y + dy), uint16(2), 0, color.RGBA{200, 200, 50, 255}}
	dx = math.Cos(o-math.Pi/6) * 40
	dy = math.Sin(o-math.Pi/6) * 40
	out <- Circle{int16(x + dx), int16(y + dy), uint16(2), 0, color.RGBA{200, 200, 50, 255}}
}

func breakFood(lm locationmanager.Location, f *food.Food, out chan Primitive) {
	ok, comb := lm.GetLocation(f)
	if !ok {
		return
	}
	x, y := comb.X, comb.Y
	out <- Circle{int16(x), int16(y), uint16(f.GetRadius()), 0, f.GetColor()}
}
