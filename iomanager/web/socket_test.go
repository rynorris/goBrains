package web

import (
	"testing"

	"code.google.com/p/go.net/websocket"
	"github.com/DiscoViking/goBrains/creature"
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
			t.Logf("Retrying...")
		} else {
			break
		}
	}
	if err != nil {
		t.Fatalf("Error connecting to websocket from %v: %v\n", origin, err)
	}

	return ws
}
