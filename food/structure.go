/*
 * Structure of food.
 *
 * These structures track the properties of creatures' food.
 */

package food

import (
	"github.com/DiscoViking/goBrains/locationmanager"
	"image/color"
)

// Food structs hold the key properties of food.
type Food struct {

	// The CollisionManager that this instance is managed by.
	cm locationmanager.Detection

	// Colour of the food.
	color color.RGBA

	// Amount of food held by this instance of the object.
	content float64
}
