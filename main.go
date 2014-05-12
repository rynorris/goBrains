package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/iomanager"
	"github.com/DiscoViking/goBrains/locationmanager"
	"github.com/banthar/Go-SDL/sdl"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var (
	drawing = true
	running = true
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

	data := make(chan []entity.Entity)
	done := make(chan struct{})
	event := make(chan sdl.Event)
	defer close(event)
	defer close(data)

	events.Global.Register(events.TERMINATE,
		func(e events.Event) { running = false; close(done) })

	lm := locationmanager.New()

	entities := make([]entity.Entity, 0, 100)
	entities = append(entities, food.New(lm, 1000))

	iomanager.Start(data, done, event)

	go func() {
		for e := range event {
			events.Handle(e)
		}

	}()

	for running {
		if drawing {
			data <- entities
			<-done
		}
		time.Sleep(12 * time.Millisecond)
	}
}
