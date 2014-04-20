/*
 * TestBrain testing.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/genetics"
	"testing"
)

// A quick kick-the-tires test for integrity.
func TestTestBrain(t *testing.T) {
	tb := newTestBrain()
	d := genetics.NewDna()
	tb.AddInputNode(brain.NewNode())
	tb.AddOutput(brain.NewNode())
	tb.Work()
	tb.Charge(0.0)
	tb.Restore(d)
}
