/*
 * Interfaces for the entity package.
 *
 * Interfaces required for entities to behave as expected.
 */

// Some witty description here.
package entity

// Entity defines the methods an entity in the environment must expose.
type Entity interface {

	// Get the size of the entity.
	getRadius() float64
}
