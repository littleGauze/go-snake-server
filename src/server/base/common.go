package base

type TokenInfo struct {
	Token   string
	Expires int
}

type User struct {
	TokenInfo
	Name string
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	NONE
)

type GameKey int

const (
	KEY_UP    GameKey = 38
	KEY_DOWN  GameKey = 40
	KEY_LEFT  GameKey = 37
	KEY_RIGHT GameKey = 39
	SPACEBAR  GameKey = 32
	W         GameKey = 87
	S         GameKey = 83
	A         GameKey = 65
	D         GameKey = 68
)
