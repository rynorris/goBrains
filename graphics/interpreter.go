/*
 * The Interpreter takes entities and breaks them down into component shapes
 * to be drawn by the lower-level graphics components.
 *
 */

package graphics

import (
	"image/color"
	"math"

	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/iomanager"
)

func Interpret(in chan iomanager.DrawSpec, out chan Primitive) {
	defer close(out)
	for spec := range in {
		switch spec.E.(type) {
		case *creature.Creature:
			breakCreature(spec, out)
		case *food.Food:
			breakFood(spec, out)
		default:
			breakEntity(spec, out)
		}
	}
}

func breakEntity(spec iomanager.DrawSpec, out chan Primitive) {
	x, y := spec.Loc.X, spec.Loc.Y

	out <- Circle{int16(x), int16(y), uint16(spec.E.Radius()), 0, spec.E.Color()}
}

func breakCreature(spec iomanager.DrawSpec, out chan Primitive) {
	var dx, dy float64

	x, y, o := spec.Loc.X, spec.Loc.Y, spec.Loc.Orient
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
	col := spec.E.Color()
	out <- Circle{int16(x), int16(y), uint16(8), 0, col}
	dx = cosO * 6
	dy = sinO * 6
	out <- Circle{int16(x - dx), int16(y - dy), uint16(6), 0, col}
	dx = cosO * 10
	dy = sinO * 10
	out <- Circle{int16(x - dx), int16(y - dy), uint16(4), 0, col}

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

func breakFood(spec iomanager.DrawSpec, out chan Primitive) {
	x, y := spec.Loc.X, spec.Loc.Y
	out <- Circle{int16(x), int16(y), uint16(spec.E.Radius()), 0, spec.E.Color()}
}
