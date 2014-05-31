package locationmanager

import "math"

const (
	zone_width  = 100
	zone_height = 100
)

type spacialZone []locatable

func (lm *LocationManager) spacialHash(c coord) (int, int) {
	x := int(math.Floor(c.locX / zone_width))
	y := int(math.Floor(c.locY / zone_height))

	return x, y
}

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

func (lm *LocationManager) findZones(hb locatable) []*spacialZone {
	zones := make([]*spacialZone, 0, 0)
	for _, coord := range hb.boundingBox() {
		have := false
		zone := lm.findZone(coord)
		for _, z := range zones {
			if z == zone {
				have = true
				break
			}
		}
		if !have {
			zones = append(zones, lm.findZone(coord))
		}
	}

	return zones
}

func (lm *LocationManager) resetZones() {
	zonesx := int(math.Ceil(TANKSIZEX / zone_width))
	zonesy := int(math.Ceil(TANKSIZEY / zone_height))

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

func (lm *LocationManager) addToZones(hb locatable) {
	zones := lm.findZones(hb)

	for _, z := range zones {
		*z = append(*z, hb)
	}
}

func (lm *LocationManager) possibleCollisions(hb locatable) []locatable {
	zones := lm.findZones(hb)

	num := 0
	for _, z := range zones {
		num += len(*z)
	}
	possibles := make([]locatable, 0, num)

	for _, z := range zones {
		possibles = append(possibles, *z...)
	}

	return possibles
}

func (lm *LocationManager) removeFromZones(hb locatable) {
	zones := lm.findZones(hb)

	for _, z := range zones {
		z.remove(hb)
	}
}

func (z *spacialZone) remove(hb locatable) {
	for i := 0; i < len(*z); i++ {
		if (*z)[i] == hb {
			// OMG MAGIC
			(*z)[i], (*z)[len(*z)-1], *z = (*z)[len(*z)-1], nil, (*z)[:len(*z)-1]
		}
	}
}
