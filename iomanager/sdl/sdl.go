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

package sdl

import (
	"fmt"

	"github.com/DiscoViking/goBrains/graphics"
	"github.com/DiscoViking/goBrains/iomanager"

	"github.com/banthar/Go-SDL/sdl"
)

const (
	WIDTH  = 800
	HEIGHT = 800
)

// Starts the graphics engine up.
// Registers it with the IoManager.
func Start(io iomanager.Manager) {
	data := make(chan []iomanager.DrawSpec, 1)
	go mainLoop(data, io)
	io.Add(iomanager.SDL, data)
}

func mainLoop(data chan []iomanager.DrawSpec, io iomanager.Manager) {
	// Initialise SDL
	fmt.Printf("Initialising SDL.")
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	// Ensure that SDL will exit gracefully when we're done.
	defer sdl.Quit()

	// Create the screen surface.
	fmt.Printf("Creating screen\n")
	screen := sdl.SetVideoMode(WIDTH, HEIGHT, 32, sdl.RESIZABLE|sdl.DOUBLEBUF|sdl.SWSURFACE)
	canvas := sdl.CreateRGBSurface(0, WIDTH, HEIGHT, 32, 0, 0, 0, 0)
	background := sdl.MapRGB(sdl.GetVideoInfo().Vfmt, 20, 20, 20)

	fmt.Printf("Entering main loop\n")
	// Main drawing loop

	// We loop every time we are passed in an array of entities to draw.
	for entities := range data {
		// Draw background.
		canvas.FillRect(&sdl.Rect{0, 0, WIDTH, HEIGHT}, background)

		// Construct the graphics pipeline.
		interpret := make(chan iomanager.DrawSpec)
		draw := make(chan graphics.Primitive)
		drawn := make(chan struct{})

		// Set off the interpreter and artist goroutines.
		go graphics.Interpret(interpret, draw)
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
		// event queue, convert them to game events, and let iomanager deal with them.
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			if ev := convert(e); ev != nil {
				io.Handle(ev)
			}
		}

		// Now wait to be told all drawing is complete before blitting to the screen
		<-drawn

		// Finally flip the surface to paint to the screen.
		screen.Blit(nil, canvas, nil)
		screen.Flip()
	}
}
