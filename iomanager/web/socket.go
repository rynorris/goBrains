package web

import (
	"log"
	"net/http"
	"time"

	"code.google.com/p/go.net/websocket"

	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/iomanager"
)

func Start(io iomanager.Manager, port string) {
	sockets := map[chan string]struct{}{}

	// The function which handles sending messages down the sockets.
	handler := func(ws *websocket.Conn) {
		log.Printf("Established connection to %v\n", ws.Request().RemoteAddr)
		in := make(chan string)
		sockets[in] = struct{}{}
		go receiveLoop(ws)
	loop:
		for s := range in {
			err := websocket.Message.Send(ws, s)
			if err != nil {
				log.Printf("Lost connection to %v\n", ws.Request().RemoteAddr)
				ws.Close()
				delete(sockets, in)
				break loop
			}
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
	go sendLoop(data, sockets)
	io.Add(iomanager.WEB, data)
}

// The loop which listens for data from IO, and sends it down the sockets.
func sendLoop(in chan []iomanager.DrawSpec, sockets map[chan string]struct{}) {
	// Use this timer to limit our sending to 30fps.
	timer := time.Tick(32 * time.Millisecond)

	for data := range in {
		json := marshal(data)
		for ws, _ := range sockets {
			select {
			case ws <- json:
			default:
			}
		}
		<-timer
	}
}

// The loop which listens for incoming events on a socket and sends them to the event handler.
func receiveLoop(ws *websocket.Conn) {
	e := event{}
	for {
		err := websocket.JSON.Receive(ws, &e)
		if err != nil {
			break
		}

		log.Printf("Received keypress %v from %v\n", e, ws.Request().RemoteAddr)
		if ev := convert(e); ev != nil {
			events.Global.Broadcast(ev)
		}
	}
}
