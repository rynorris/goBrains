/*
 * TestBrain testing.
 */

package creature

import (
	"testing"

	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/genetics"
)

// A quick kick-the-tires test for integrity.
func TestTestBrain(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	tb := newTestBrain()
	d := genetics.NewDna()
	tb.AddInputNode(brain.NewNode())
	tb.AddOutput(brain.NewNode())
	tb.GenesNeeded()
	tb.Work()
	tb.Charge(0.0)
	tb.Restore(d)
}
