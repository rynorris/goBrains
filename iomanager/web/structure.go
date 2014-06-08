package web

import (
	"image/color"
)

type entitySpec struct {
	Type string
	X    int
	Y    int
}

type creatureSpec struct {
	entitySpec

	Colour color.RGBA
	Angle  float64
}

type foodSpec struct {
	entitySpec

	Size int
}
