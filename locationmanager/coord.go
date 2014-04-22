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
