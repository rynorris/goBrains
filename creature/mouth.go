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
func (mt *Mouth) detect() {
	blips := mt.host.lm.GetCollisions(mt.location, mt.host)

	// Attempt to consume all entities at this position.
	for _, blip := range blips {
		a := blip.Consume()
		mt.host.vitality += a
		mt.node.Charge(a)
	}
}

// Add a new mouth to a creature.
// This is at the front of the creature.
func (host *creature) AddMouth() *Mouth {
	node := brain.NewNode()
	host.brain.AddInputNode(node)

	input := inputStruct{
		putStruct: putStruct{host: host},
		node:      node,
		location:  locationmanager.CoordDelta{host.Radius(), 0.0},
	}
	m := &Mouth{input}

	host.inputs = append(host.inputs, m)
	return m
}
