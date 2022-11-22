package internal

type Position struct {
	X int
	Y int
}

func (p Position) NewPosition(x, y int) Position {
	return Position{x, y}
}

func (p Position) ClonePosition() Position {
	return Position{p.X, p.Y}
}
