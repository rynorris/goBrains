package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/entitymanager"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/iomanager"
	"github.com/DiscoViking/goBrains/iomanager/web"
	"github.com/DiscoViking/goBrains/stats"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var headless = flag.Bool("headless", false, "run in headless mode")
var port = flag.Int("port", 9999, "port to listen on for web connections")
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

		// Run for 30 seconds and stop.
		stop := time.After(30 * time.Second)
		go func() {
			<-stop
			running = false
		}()

	}

	rand.Seed(time.Now().UnixNano())

	config.Load("config.gcfg")

	em := entitymanager.New()
	em.Reset()

	io := iomanager.New(em.LocationManager())
	defer io.Shutdown()

	if !*headless {
		rateLimit = true
	}

	drawTimer := time.Tick(16 * time.Millisecond)
	ticktime := 16 * time.Millisecond
	tickTimer := time.Tick(ticktime)

	web.Start(io, strconv.Itoa(*port))
	stats.Start()

	events.Global.Register(events.TERMINATE,
		func(e events.Event) { running = false })
	events.Global.Register(events.TOGGLE_DRAW,
		func(e events.Event) { drawing = !drawing })
	events.Global.Register(events.TOGGLE_FRAME_LIMIT,
		func(e events.Event) { rateLimit = !rateLimit })
	events.Global.Register(events.SPEED_UP,
		func(e events.Event) {
			ticktime /= 2
			tickTimer = time.Tick(ticktime)
		})
	events.Global.Register(events.SPEED_DOWN,
		func(e events.Event) {
			ticktime *= 2
			tickTimer = time.Tick(ticktime)
		})

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
			oldest, _ := stats.Global.Oldest()
			log.Printf("Stats:\n FPS: %v\n Population: %v\n Average: %v\n Oldest: %v\n\n",
				frames/2, stats.Global.Population(), stats.Global.AverageAge(), oldest)
			frames = 0
		}
		em.Spin()
		if rateLimit {
			<-tickTimer
		}
		select {
		case <-drawTimer:
			drawFunc()
		default:
		}
	}
}
