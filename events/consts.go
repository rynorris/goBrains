package events

type EventType int

const (
	NONE = iota
	TERMINATE
	TOGGLE_DRAW
	SELECT
	TOGGLE_FRAME_LIMIT
	SPEED_UP
	SPEED_DOWN
)

var (
	Global Handler = NewHandler()
)
