package iomanager

import (
	"testing"

	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/locationmanager"
)

func TestLocate(t *testing.T) {
	lm := locationmanager.NewLocationManager(200, 200)

	// Add an entity to LM.
	e := &entity.TestEntity{}
	loc := locationmanager.Combination{45, 23, 1.5}
	lm.AddEntAtLocation(e, loc)

	// Try to find it. This should succeed.
	out, ok := Locate(lm, e)

	if !ok {
		t.Errorf("Failed to locate locatable entity.")
	}

	exp := DrawSpec{e, loc}
	if out != exp {
		t.Errorf("Found wrong location. Exp: %v, Got: %v", loc, out)
	}

	// Made a new entity, but don't add it to LM.
	e = &entity.TestEntity{}

	// The lookup should fail this time.
	out, ok = Locate(lm, e)

	if ok {
		t.Errorf("Locateed unlocatable entity at: %v", out)
	}
}
