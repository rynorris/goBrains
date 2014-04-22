/*
 * Booster structures.
 *
 * This covers the velocity output behaviours that result in a change in location for a creature.
 */

package creature

// Fixed values.
const (
	// Booster types.
	BoosterLinear  = 1
	BoosterAngular = 2
)

// Outputs are chargeable.  This means they accept accept charge from nodes in the brain.
func (b *booster) Charge(strength float64) {
	b.charge += strength
}

// Outputs are workers.  This means that the brain will trigger them during processing to perform their actions.
func (b *booster) Work() {

	if b.btype == BoosterLinear {
		b.host.movement.move += b.charge
	} else if b.btype == BoosterAngular {
		b.host.movement.rotate += b.charge
	}

	// Reset charge now it has been used.
	b.charge = 0
}

// Initialize a new generic booster object.
func (host *Creature) newGenBooster(btype int) *booster {

	newBoost := booster{
		outputStruct: outputStruct{
			putStruct: putStruct{host: host},
			charge:    0,
		},
		btype: btype,
	}
	b := &newBoost

	// Link the output to the hosts' brain.
	host.brain.AddOutput(b)

	return b
}

// Add a standard set of boosters to a host; one angular and one linear.
func (host *Creature) AddBoosters() (*booster, *booster) {
	l := host.newGenBooster(BoosterLinear)
	a := host.newGenBooster(BoosterAngular)
	return l, a
}
