/*
 * The GraphicsManager deals with initialising and scheduling all graphics events.
 *
 * It will initialise the SDL module, and then arrange the piping of all Entities through
 * the graphics pipeline.
 *
 * The current graphics pipeline is:
 * External -> Interpreter -> Artist -> Screen
 *
 * The Interpreter breaks entities down into primitive shapes, which the artist draws
 * to the screen.
 */

package graphics

import "github.com/discoviking/goBrains/entity"

import "github.com/banthar/Go-SDL/sdl"

// Starts the graphics engine up.
// Once called, pass a slice of Entities into the passed in channel
// once per frame for drawing.
// Then wait on the done channel for drawing to finish before continuing.
func Start(data chan []entity.Entity, done chan struct{}) {

	// Initialise SDL
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	// Construct the graphics pipeline.
	interpret := make(chan entity.Entity)
	draw := make(chan Primitive)

	// Create the screen surface.
	screen := sdl.SetVideoMode(640, 480, 32, sdl.RESIZABLE)

	// Set off the interpreter and artist goroutines.
	go Interpret(interpret, draw)
	go Draw(draw, screen)

	// Main drawing loop
	for entities := range data {
		for _, e := range entities {
			// Send entities to the interpreter
			interpret <- e
		}

		// This is pretty dangerous, since we're continuing without waiting for the artist
		// to finish drawing the final entity to screen.
		// I think that the worst that will happen is that one entity will be drawn a single
		// frame behind the others, which is no great loss. However if necessary this can
		// be fixed, it just requires a bit more synchronization rubbish and slightly ruins
		// the cleanliness of this pipeline design.
		screen.Flip()
		done <- struct{}{}
	}
}
