package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"code.google.com/p/go.net/websocket"

	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/iomanager"
)

func Start(io iomanager.Manager, port string) {
	sockets := make([]chan string, 0, 0)

	// The function which handles sending messages down the sockets.
	handler := func(ws *websocket.Conn) {
		in := make(chan string)
		sockets = append(sockets, in)
		go receiveLoop(ws)
		for s := range in {
			websocket.Message.Send(ws, s)
		}
	}

	// Kick off the http server handling socket creation requests.
	go func() {
		http.HandleFunc("/data",
			func(w http.ResponseWriter, req *http.Request) {
				s := websocket.Server{Handler: websocket.Handler(handler)}
				s.ServeHTTP(w, req)
			})

		http.HandleFunc("/tank",
			func(w http.ResponseWriter, req *http.Request) {
				http.ServeFile(w, req, "iomanager/web/tank.html")
			})
		log.Printf("Listening on port %v.\n", port)
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Printf("Failed to start HTTP server on port %v: %v.\n", port, err)
		}
	}()

	// The channel by which iomanager will send us stuff
	data := make(chan []iomanager.DrawSpec, 0)
	go sendLoop(data, &sockets)
	io.Add(iomanager.WEB, data)
}

// The loop which listens for data from IO, and sends it down the sockets.
func sendLoop(in chan []iomanager.DrawSpec, sockets *[]chan string) {
	// Use this timer to limit our sending to 30fps.
	timer := time.Tick(32 * time.Millisecond)

	for data := range in {
		json := marshal(data)
		for _, ws := range *sockets {
			ws <- json
		}
		<-timer
	}
}

// The loop which listens for incoming events on a socket and sends them to the event handler.
func receiveLoop(ws *websocket.Conn) {
	e := event{}
	for {
		websocket.JSON.Receive(ws, &e)
		fmt.Println("Got: %v", e)
		if ev := convert(e); ev != nil {
			events.Global.Broadcast(ev)
		}
	}
}
