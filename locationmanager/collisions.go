package locationmanager

import "math"

const (
	zone_width  = 150
	zone_height = 150
)

type spacialZone []locatable

// Converts a location into two integers representing which spacial
// zone it lies in.
func (lm *LocationManager) spacialHash(c coord) (int, int) {
	x := int(math.Floor(c.locX / zone_width))
	y := int(math.Floor(c.locY / zone_height))

	return x, y
}

// findZone returns the spacial zone the given point lies in.
func (lm *LocationManager) findZone(c coord) *spacialZone {
	x, y := lm.spacialHash(c)
	if x < 0 {
		x = 0
	}
	if x >= len(lm.spacialZones) {
		x = len(lm.spacialZones) - 1
	}
	if y < 0 {
		y = 0
	}
	if y >= len(lm.spacialZones[x]) {
		y = len(lm.spacialZones[x]) - 1
	}
	return &lm.spacialZones[x][y]
}

// findZones adds all the zones a hitbox lies in to it's zone store.
func (lm *LocationManager) findZones(hb locatable) {
	hb.clearZones()
	for _, coord := range hb.boundingBox() {
		have := false
		zone := lm.findZone(coord)
		for _, z := range hb.zones() {
			if z == zone {
				have = true
				break
			}
		}
		if !have {
			hb.addZone(zone)
		}
	}
}

// resetZones clears all hitboxes out of all zones.
// Note it does not remove all zones from all hitboxes. However this
// shouldn't be an issue since this is only called once at start of day.
func (lm *LocationManager) resetZones() {
	zonesx := int(math.Ceil(lm.maxPoint.locX / zone_width))
	zonesy := int(math.Ceil(lm.maxPoint.locY / zone_height))

	lm.spacialZones = make([][]spacialZone, 0, zonesx)

	for i := 0; i < zonesx; i++ {
		col := make([]spacialZone, 0, zonesy)
		for j := 0; j < zonesy; j++ {
			zone := make(spacialZone, 0, 20)
			col = append(col, zone)
		}
		lm.spacialZones = append(lm.spacialZones, col)
	}
}

// addToZones put a hitbox in all the zones it lies in, and adds
// all those zones to it's zone store.
func (lm *LocationManager) addToZones(hb locatable) {
	lm.findZones(hb)

	for _, z := range hb.zones() {
		*z = append(*z, hb)
	}
}

/* Function commented out since not currently used.

// possibleCollisions returns a slice of locatables which may possibly
// collide with the given one.
func (lm *LocationManager) possibleCollisions(hb locatable) []locatable {
	num := 0
	for _, z := range hb.zones() {
		num += len(*z)
	}
	possibles := make([]locatable, 0, num)

	for _, z := range hb.zones() {
		for _, h := range *z {
			have := false
			for _, a := range possibles {
				if a == h {
					have = true
					break
				}
			}
			if !have {
				possibles = append(possibles, h)
			}
		}
	}

	return possibles
}
*/

// removeFromZones removes a hitbox from all zones it was in,
// and clears out it's zone store.
func (lm *LocationManager) removeFromZones(hb locatable) {
	for _, z := range hb.zones() {
		z.remove(hb)
	}
	hb.clearZones()
}

// remove takes a hitbox out of the given zone.
func (z *spacialZone) remove(hb locatable) {
	for i := 0; i < len(*z); i++ {
		if (*z)[i] == hb {
			// OMG MAGIC
			(*z)[i], (*z)[len(*z)-1], *z = (*z)[len(*z)-1], nil, (*z)[:len(*z)-1]
		}
	}
}
