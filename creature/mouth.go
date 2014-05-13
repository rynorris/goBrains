/*
 * Mouth.
 *
 * Food consumption organ for a creature.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/locationmanager"
)

// Yum.
func (mt *mouth) detect() {
	blips := mt.host.lm.GetCollisions(mt.location, mt.host)

	// Attempt to consume all entities at this position.
	for _, blip := range blips {
		mt.host.vitality += blip.Consume()
		mt.node.Charge(0.2)
	}
}

// Add a new mouth to a creature.
// This is at the front of the creature.
func (host *Creature) AddMouth() *mouth {
	node := brain.NewNode()
	host.brain.AddInputNode(node)

	input := inputStruct{
		putStruct: putStruct{host: host},
		node:      node,
		location:  locationmanager.CoordDelta{host.GetRadius(), 0.0},
	}
	m := &mouth{input}

	host.inputs = append(host.inputs, m)
	return m
}
