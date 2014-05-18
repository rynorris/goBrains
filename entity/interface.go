/*
 * Interfaces for the entity package.
 *
 * Interfaces required for entities to behave as expected.
 */

// Some witty description here.
package entity

import "image/color"

// Entity defines the methods an entity in the environment must expose.
type Entity interface {

	// Get the size of the entity.
	Radius() float64

	// Get the colour of this creature.
	Color() color.RGBA

	// Query the state of the entity.
	// Returns a boolean for if it is being torn down on this check.
	Check() bool

	// Attempt to consume the enitity.
	// Returns the amount consumed.
	Consume() float64
}
