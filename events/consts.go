package events

type EventType int

const (
	NONE               = 0
	TERMINATE          = 1
	TOGGLE_DRAW        = 2
	SELECT             = 3
	TOGGLE_FRAME_LIMIT = 4
)

var (
	Global Handler = NewHandler()
)
