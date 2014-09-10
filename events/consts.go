package events

type EventType int

const (
	NONE = iota

	// Events related to user input.
	TERMINATE
	TOGGLE_DRAW
	SELECT
	TOGGLE_FRAME_LIMIT
	SPEED_UP
	SPEED_DOWN

	// Internal events used for communication between components.
	POPULATION_STATE
	ENTITY_DESTROY
	ENTITY_CREATE
)

var (
	Global Handler = NewHandler()
)
