package locationmanager

// Tests to specifically validate the collision code.

import (
	"testing"
)

func TestFindZone(t *testing.T) {
	// Create an instance of LM which will contain exactly 4 zones.
	lm := NewLocationManager(zone_width*2, zone_height*2)

	// List of test cases
	testcases := map[coord][]int{
		// Test the corners of each zone.
		coord{0, 0}:                    []int{0, 0},
		coord{0, zone_height}:          []int{0, 1},
		coord{zone_width, 0}:           []int{1, 0},
		coord{zone_width, zone_height}: []int{1, 1},

		// The centre of each zone.
		coord{zone_width / 2, zone_height / 2}:     []int{0, 0},
		coord{zone_width * 1.5, zone_height / 2}:   []int{1, 0},
		coord{zone_width / 2, zone_height * 1.5}:   []int{0, 1},
		coord{zone_width * 1.5, zone_height * 1.5}: []int{1, 1},

		// Some off-screen points.
		coord{-10, -10}:                        []int{0, 0},
		coord{0, zone_height * 5}:              []int{0, 1},
		coord{zone_width * 5, 0}:               []int{1, 0},
		coord{zone_width * 5, zone_height * 5}: []int{1, 1},
	}

	for coord, val := range testcases {
		zone := lm.findZone(coord)
		if zone != &lm.spacialZones[val[0]][val[1]] {
			t.Errorf("Incorrect zone found for point %v.\nExpected %v=%v, got %v.\n",
				coord, val, &lm.spacialZones[val[0]][val[1]], zone)
		}
	}
}
