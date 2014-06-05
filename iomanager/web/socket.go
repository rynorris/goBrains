package web

import (
	"log"
	"net/http"

	"github.com/DiscoViking/goBrains/iomanager"
)

func Start(io iomanager.Manager) {
	sockets := make([]chan string, 0, 0)
	go sendLoop(io, &sockets)

	// The function which handles sending messages down the sockets.
	handler := func(ws *websocket.Conn) {
		in := make(chan string)
		sockets = append(sockets, in)
		for s := range in {
			websocket.Message.Send(ws, s)
		}
	}

	// Kick off the http server handling socket creation requests.
	go func() {
		http.HandleFunc("/tank",
			func(w http.ResponseWriter, req *http.Request) {
				s := websocket.Server{Handler: websocket.Handler(handler)}
				s.ServeHTTP(w, req)
			})
		err := http.ListenAndServe(":10000", nil)
		if err != nil {
			log.Println(err)
		}
	}()

	// The channel by which iomanager will send us stuff
	data := make(chan []iomanager.DrawSpec, 1)
	io.Add(iomanager.WEB, data)
}

// The loop which listens for data from IO, and sends it down the sockets.
func sendLoop(in chan []iomanager.DrawSpec, sockets *[]chan string) {
	for data := range in {
		json := marshal(data)
		for _, ws := range *sockets {
			ws <- json
		}
	}
}
