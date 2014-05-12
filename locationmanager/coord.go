/*
 * Coordiate manipulation.
 *
 * Methods for manipulating coordinates.
 */

package locationmanager

// Update a co-ordinate with a change.
func (loc *coord) update(deltaX, deltaY float64) {
	loc.locX += deltaX
	loc.locY += deltaY
}

// Limit the coordinate to within a rectangle.
func (loc *coord) limit(min *coord, max *coord) {
	if loc.locX < min.locX {
		loc.locX = min.locX
	} else if loc.locX > max.locX {
		loc.locX = max.locX
	}

	if loc.locY < min.locY {
		loc.locY = min.locY
	} else if loc.locY > max.locY {
		loc.locY = max.locY
	}
}
