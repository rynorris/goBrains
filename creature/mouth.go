/*
 * Mouth.
 *
 * Food consumption organ for a creature.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/locationmanager"
)

// Yum.
func (mt *mouth) detect() {
	blips := mt.host.lm.GetCollisions(mt.location, mt.host)

	// Attempt to consume all entities at this position.
	for _, blip := range blips {
		mt.host.vitality += blip.Consume()
	}
}

// Create a new mouth.
// This is at the front of the creature.
func newMouth(host *Creature) *mouth {
	input := inputStruct{
		putStruct: putStruct{host: host},
		node:      nil,
		location:  locationmanager.CoordDelta{host.GetRadius(), 0.0},
	}
	return &mouth{input}
}
