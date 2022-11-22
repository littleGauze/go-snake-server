package internal

type ControlKey int

const (
	START ControlKey = iota
	PAUSE
	RESET
)

type ScreenEdge int

const (
	NORTH ScreenEdge = iota
	SOUTH
	EAST
	WEST
)

type Speed int

const (
	SLOW Speed = iota
	FAST
	NORMAL
)
