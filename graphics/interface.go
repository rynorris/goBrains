/*
 * Interfaces for the graphics manager
 *
 * These interfaces provide the method by which other components communicate with GM.
 */

package graphics

type GraphicsManager interface {
	Start(width, height int)
}
