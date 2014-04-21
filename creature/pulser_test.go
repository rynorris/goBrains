/*
 * Testing for pulsers.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/locationmanager"
	"testing"
)

func TestPulsers(t *testing.T) {
	host := New(locationmanager.NewLocationManager())
	tBrain := newTestBrain()
	host.brain = tBrain
	var expected int

	// Add a pulser to the creature.
	p := host.AddPulser()

	// Each time the pulser runs a 5 detection cycles it should charge the brain once.
	// Don't run more than i == 9, as the patten is more complex due to the way nodes decay charge.
	for i := 0; i < 10; i++ {
		p.detect()
		tBrain.Work()
		expected = (i / 5)
		if tBrain.fired != expected {
			t.Errorf("Expected to fire %v, actually fired %v times.", expected, tBrain.fired)
		}
	}
}
