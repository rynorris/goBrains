/*
 * Interfaces for the brain.
 *
 * These interfaces provides the mechanisms by which elements in a neural network act upon each other.
 */

// Package brain provides a real-time neural net framework.
package brain

// Chargeable defines the family of objects that participate in a neural network.
// Objects that implement the chargable interface accept ChargeCarrier from other objects in the neural net.
type Chargeable interface {
	Charge(strength float64)
}

// Worker defines the family of objects which can be added to the neural worker queue and are scheduled for later work.
// Objects that implement the worker interface perform actions only during a work process.
type Worker interface {
	Work()
}

// ChargedWorker combines the Chargeable and Worker interfaces.
// (is there a better way to do this?)
type ChargedWorker interface {
	Charge(strength float64)
	Work()
}
