package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/entitymanager"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/iomanager/sdl"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var (
	drawing   = true
	running   = true
	rateLimit = true
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	rand.Seed(time.Now().UnixNano())

	done := make(chan struct{})
	event := make(chan events.Event)
	defer close(event)
	defer close(done)

	em := entitymanager.New()
	em.Reset()

	outputs := make([]chan []entity.Entity, 0, 1)
	data := sdl.Start(em.LocationManager(), done, event)
	outputs = append(outputs, data)
	defer close(data)

	go func() {
		for e := range event {
			events.Global.Broadcast(e)
		}

	}()

	timer := time.Tick(16 * time.Millisecond)

	events.Global.Register(events.TERMINATE,
		func(e events.Event) { running = false })
	events.Global.Register(events.TOGGLE_DRAW,
		func(e events.Event) { drawing = !drawing })
	events.Global.Register(events.TOGGLE_FRAME_LIMIT,
		func(e events.Event) { rateLimit = !rateLimit })

	drawFunc := func() {
		for _, data := range outputs {
			if drawing {
				data <- em.Entities()
			} else {
				data <- []entity.Entity{}
			}
			<-done
		}
	}

	frames := 0
	before := time.Now()

	for running {
		frames += 1
		if time.Since(before) > 2*time.Second {
			before = time.Now()
			log.Printf("FPS: %v\n", frames/2)
			frames = 0
		}
		em.Spin()
		if rateLimit {
			<-timer
			drawFunc()
		} else {
			select {
			case <-timer:
				drawFunc()
			default:
			}
		}
	}
}
