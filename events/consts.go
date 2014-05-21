package events

type EventType int

const (
	NONE = iota
	TERMINATE
	TOGGLE_DRAW
	SELECT
	TOGGLE_FRAME_LIMIT
)

var (
	Global Handler = NewHandler()
)
