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
	"github.com/banthar/Go-SDL/sdl"
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

	cfg, err := config.Load("config.gcfg")
	if err != nil {
		log.Fatal(err)
	}

	loadModules(cfg)

	data := make(chan []entity.Entity)
	done := make(chan struct{})
	event := make(chan sdl.Event)
	defer close(event)
	defer close(data)
	defer close(done)

	em := entitymanager.New()
	em.Reset()

	iomanager.Start(em.LocationManager(), data, done, event)

	go func() {
		for e := range event {
			events.Handle(e)
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
		if drawing {
			data <- em.Entities()
		} else {
			data <- []entity.Entity{}
		}
		<-done
	}

	for running {
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

// Loads config into all modules which require it.
func loadModules(cfg *config.Config) {
	entitymanager.LoadConfig(cfg)
	iomanager.LoadConfig(cfg)
}
