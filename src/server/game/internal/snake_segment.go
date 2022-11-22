package internal

import (
	"math"
)

var Colors = []string{"#FF0000", "#FF9966",
	"#FFFA66", "#66FF66",
	"#66FFFD", "#6699FF",
	"#7966FF", "#F366FF"}

type SnakeSegment struct {
	GameObject
	colorIndex int `json:"colorIndex"`
}

func (ss *SnakeSegment) Init(pos Position) {
	ss.Clazz = "SnakeSegment"
	ss.colorIndex = -1
	ss.Position = pos
}

func (ss *SnakeSegment) GetColor() string {
	ss.colorIndex++
	return Colors[int(math.Mod(float64(ss.colorIndex), float64(len(Colors))))]
}

func (ss *SnakeSegment) HandleCollision(snake *Snake) {
	snake.Die()
}
