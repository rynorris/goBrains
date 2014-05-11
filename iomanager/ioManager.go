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

import "fmt"
import "github.com/DiscoViking/goBrains/entity"
import "github.com/DiscoViking/goBrains/graphics"
import "github.com/DiscoViking/goBrains/events"

import "github.com/banthar/Go-SDL/sdl"

const (
	WIDTH  = 640
	HEIGHT = 480
)

// Starts the graphics engine up.
// Once called, pass a slice of Entities into the passed in channel
// once per frame for drawing.
// Then wait on the done channel for drawing to finish before continuing.
func Start(data chan []entity.Entity, done chan struct{}, handle chan events.Event) {

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
	background := sdl.MapRGB(sdl.GetVideoInfo().Vfmt, 200, 200, 200)

	fmt.Printf("Entering main loop\n")
	// Main drawing loop
	time := uint32(0)
	frame := 0
	// We loop every time we are passed in an array of entities to draw.
	for entities := range data {
		frame = (frame + 1) % 100
		if frame == 0 {
			newTime := sdl.GetTicks()
			dt := newTime - time
			time = newTime
			fps := 100000 / float32(dt)
			fmt.Printf("Dt: %v, FPS: %v\n", dt, fps)
		}

		canvas.FillRect(&sdl.Rect{0, 0, WIDTH, HEIGHT}, background)

		// Construct the graphics pipeline.
		interpret := make(chan entity.Entity)
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
		// event queue and send them to the event manager.
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			ev := events.Convert(e)
			if ev.GetType() != events.NONE {
				handle <- ev
			}
		}

		// Now wait to be told all drawing is complete before blitting to the screen
		<-drawn

		// Finally flip the surface to paint to the screen.
		screen.Blit(nil, canvas, nil)
		screen.Flip()
		done <- struct{}{}
	}
}
