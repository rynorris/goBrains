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

package iomanager

import (
	"fmt"

	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/graphics"
	"github.com/DiscoViking/goBrains/locationmanager"

	"github.com/banthar/Go-SDL/sdl"
)

// Starts the graphics engine up.
// Once called, pass a slice of Entities into the passed in channel
// once per frame for drawing.
// Then wait on the done channel for drawing to finish before continuing.
func Start(lm locationmanager.Location, data chan []entity.Entity, done chan struct{}, handle chan sdl.Event) {
	go mainLoop(lm, data, done, handle)
}

func mainLoop(lm locationmanager.Location, data chan []entity.Entity, done chan struct{}, handle chan sdl.Event) {

	// Initialise SDL
	fmt.Printf("Initialising SDL.")
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	// Create the screen surface.
	fmt.Printf("Creating screen\n")
	screen := sdl.SetVideoMode(width, height, 32, sdl.RESIZABLE|sdl.DOUBLEBUF|sdl.SWSURFACE)
	canvas := sdl.CreateRGBSurface(0, width, height, 32, 0, 0, 0, 0)
	background := sdl.MapRGB(sdl.GetVideoInfo().Vfmt, 20, 20, 20)

	fmt.Printf("Entering main loop\n")
	// Main drawing loop

	// We loop every time we are passed in an array of entities to draw.
	for entities := range data {

		canvas.FillRect(&sdl.Rect{0, 0, uint16(width), uint16(height)}, background)

		// Construct the graphics pipeline.
		interpret := make(chan entity.Entity)
		draw := make(chan graphics.Primitive)
		drawn := make(chan struct{})

		// Set off the interpreter and artist goroutines.
		go graphics.Interpret(lm, interpret, draw)
		go graphics.Draw(draw, canvas, drawn)

		for _, e := range entities {
			// Send entities to the interpreter
			interpret <- e
		}

		// Close the interpret pipe. This will cause the interpreter to finish.
		// It will then close the drawing pipe, causing the artist to finish.
		// He will then signal to us that he is done via the drawn pipe.
		close(interpret)

		// Whilst the drawing is potentially still going, pull any events off the SDL
		// event queue and send them to the event manager.
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			handle <- e
		}

		// Now wait to be told all drawing is complete before blitting to the screen
		<-drawn

		// Finally flip the surface to paint to the screen.
		screen.Blit(nil, canvas, nil)
		screen.Flip()
		done <- struct{}{}
	}
}
