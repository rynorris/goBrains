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
func newGenBooster(host *Creature, btype int) *booster {

	newBoost := booster{
		putStruct: putStruct{host: host},
		charge:    0,
		btype:     btype,
	}

	// Link the output to the hosts' brain.
	host.brain.AddOutput(&newBoost)

	return &newBoost
}

// Initialize a new linear booster.
func newLinearBooster(host *Creature) *booster {
	return newGenBooster(host, BoosterLinear)
}

// Initialize a new angular booster.
func newAngularBooster(host *Creature) *booster {
	return newGenBooster(host, BoosterAngular)
}
