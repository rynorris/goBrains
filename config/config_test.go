package config

import (
	"reflect"
	"testing"
)

var (
	expected = &Config{
		General: struct {
			ScreenWidth  int
			ScreenHeight int
		}{800, 800},
		Entity: struct {
			InitialCreatures int
			InitialFood      int
			BreedingRate     int
			FoodSpawnRate    int
			FoodSize         float64
		}{40, 50, 150, 200, 1000},
		Brain: struct {
			ChargeDecayRate     float64
			SynapseMaxCharge    float64
			SynapseOutputScale  float64
			NodeFiringThreshold float64
			NodeFiringStrength  float64
		}{0.02, 1.0, 0.1, 1.0, 0.8},
		Genetics: struct {
			MutationRate int
		}{500},
	}
)

func TestConfig(t *testing.T) {
	cfg, err := Load("test/test_config.gcfg")
	if err != nil {
		t.Fatalf("Failed to load test config: %v.\n", err)
	}

	if !reflect.DeepEqual(cfg, expected) {
		t.Errorf("Loaded config did not match expected.\nGot: %v\nExpected: %v\n", cfg, expected)
	}
}
