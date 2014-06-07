package web

import (
	"testing"
	"time"

	"code.google.com/p/go.net/websocket"
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/iomanager"
	"github.com/DiscoViking/goBrains/locationmanager"
)

func TestSocketFV(t *testing.T) {
	// Set up the inputs/outputs for this test.
	lm := locationmanager.NewLocationManager(400, 400)
	io := iomanager.New(lm)

	c := creature.NewSimple(lm)
	input := []iomanager.DrawSpec{
		{c, locationmanager.Combination{20, 30, 0.5}},
	}
	expected := `{"scene":{"entities":[{"Type":"creature","X":20,"Y":30,"Colour":{"R":200,"G":50,"B":50,"A":255},"Angle":0.5}]}}`

	// Start the websocket output.
	Start(io, "9999")

	// Create a client.
	ws := newClient(t, "http://client1.localhost")

	io.Out[iomanager.WEB] <- input
	receiveAndCheck(t, ws, expected)

	// Now test with multiple clients.
	// Create 2 more clients.
	ws2 := newClient(t, "http://client2.localhost")
	ws3 := newClient(t, "http://client3.localhost")

	// Check they all receive the data
	io.Out[iomanager.WEB] <- input
	receiveAndCheck(t, ws, expected)
	receiveAndCheck(t, ws2, expected)
	receiveAndCheck(t, ws3, expected)

	// Check that input works ok.
	// Register us for events.
	received := make(chan struct{})
	timeout := time.After(1 * time.Second)
	events.Global.Register(events.TOGGLE_FRAME_LIMIT, func(e events.Event) {
		t.Logf("%v", e)
		received <- struct{}{}
	})

	websocket.Message.Send(ws, `{"Type":"Key","Key":"Y"}`)
	select {
	case <-received:
		t.Errorf("Received toggle frame limit event on Y keypress.")
	case <-timeout:
		t.Logf("Didn't receive toggle frame limit event on Y keypress.")
	}

	timeout = time.After(1 * time.Second)
	websocket.Message.Send(ws, `{"Type":"Key","Key":"Z"}`)
	select {
	case <-received:
		t.Logf("Received toggle frame limit event on Z keypress.")
	case <-timeout:
		t.Errorf("Didn't receive toggle frame limit event on Z keypress.")
	}
}

func receiveAndCheck(t *testing.T, ws *websocket.Conn, expected string) {
	var received string
	websocket.Message.Receive(ws, &received)

	if received != expected {
		t.Errorf("Expected: %v\nGot: %v\n", expected, received)
	}
}

func newClient(t *testing.T, origin string) *websocket.Conn {
	var ws *websocket.Conn
	var err error
	for i := 0; i < 3; i++ {
		ws, err = websocket.Dial("ws://localhost:9999/data", "", origin)
		if err != nil {
			t.Logf("Failed to connect to websocket from %v: %v\n", origin, err)
			if i < 2 {
				t.Logf("Retrying...")
			}
		} else {
			break
		}
	}
	if err != nil {
		t.Fatalf("Error connecting to websocket from %v: %v\n", origin, err)
	} else {
		t.Logf("Successfully connected client at %v", origin)
	}

	return ws
}
