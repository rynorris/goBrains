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
        default:
            out <- breakEntity(e)
        }
    }
}

func breakEntity(e entity.Entity) Primitive {
    return Circle{20, 20, uint16(e.GetRadius()), color.Black}
}

func breakFood(f food.Food) Primitive {
    return Circle{20, 20, uint16(f.GetRadius()), color.RGBA{50, 200, 50, 255}}
}
