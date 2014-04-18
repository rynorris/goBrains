/*
 * Creature interfaces.
 *
 * These interfaces are those internally and externally applicable to a creature.
 */

package creature

// Interface exposed by input objects.
type input interface {

	// Activate an input, supplying environmental information to the brain.
	detect()
}
