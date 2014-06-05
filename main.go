package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/entitymanager"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/iomanager"
	"github.com/DiscoViking/goBrains/iomanager/sdl"
	"github.com/DiscoViking/goBrains/iomanager/web"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var headless = flag.Bool("headless", false, "run in headless mode")
var (
	drawing   = true
	running   = true
	rateLimit = false
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

	config.Load("config.gcfg")

	em := entitymanager.New()
	em.Reset()

	io := iomanager.New(em.LocationManager())
	defer io.Shutdown()

	if !*headless {
		sdl.Start(io)
		rateLimit = true
	}

	web.Start(io)

	timer := time.Tick(8 * time.Millisecond)

	events.Global.Register(events.TERMINATE,
		func(e events.Event) { running = false })
	events.Global.Register(events.TOGGLE_DRAW,
		func(e events.Event) { drawing = !drawing })
	events.Global.Register(events.TOGGLE_FRAME_LIMIT,
		func(e events.Event) { rateLimit = !rateLimit })

	drawFunc := func() {
		if drawing {
			io.Distribute(em.Entities())
		} else {
			io.Distribute([]entity.Entity{})
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
