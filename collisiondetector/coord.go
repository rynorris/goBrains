/*
 * Coordiate manipulation.
 *
 * Methods for manipulating coordinates.
 */

package collisiondetector

// Update a co-ordinate with a change.
func (loc coord) update(delta CoordDelta) {
	loc.locX += delta.deltaX
	loc.locY += delta.deltaY
}
