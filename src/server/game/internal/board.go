package internal

import (
	"encoding/json"
	"math/rand"
	"server/base"
	"time"

	"github.com/name5566/leaf/log"
)

const (
	BOARD_WIDTH  int    = 600
	BOARD_HEIGHT int    = 400
	BLOCK_SIZE   int    = 8
	BG_COLOR     string = "#fff"
	GRID_COLOR   string = "#001F5C"
)

type Board struct {
	width  int
	height int

	grid [][]*Drawable
}

func (b *Board) PlaceObject(obj Drawable, pos Position) {
	b.grid[pos.X][pos.Y] = &obj
	obj.SetPosition(pos.ClonePosition())
}

func (b *Board) RemoveObject(pos Position) {
	b.grid[pos.X][pos.Y] = nil
}

func (b *Board) MoveObject(obj Drawable, pos Position) {
	b.RemoveObject(obj.GetPosition())
	b.PlaceObject(obj, pos)
}

func (b *Board) PlaceAtRandom(obj Drawable) {
	p := b.getRandomPosition()
	b.PlaceObject(obj, p)
}

func (b *Board) getRandomPosition() Position {
	var pos *Position
	for {
		if pos != nil {
			return *pos
		}
		x := GetRandom(b.width)
		y := GetRandom(b.height)
		if b.grid[x][y] == nil {
			pos = &Position{X: x, Y: y}
		}
	}
}

func (b *Board) GetRandomPositionAndDirection() (Position, base.Direction) {
	dirs := []base.Direction{base.UP, base.DOWN, base.LEFT, base.RIGHT}
	dir := dirs[GetRandom(len(dirs))]

	var pos Position
	switch dir {
	case base.UP:
		pos = Position{X: GetRandom(b.width), Y: b.height - 1}
	case base.DOWN:
		pos = Position{X: GetRandom(b.width), Y: 0}
	case base.LEFT:
		pos = Position{X: b.width - 1, Y: GetRandom(b.height)}
	case base.RIGHT:
		pos = Position{X: 0, Y: GetRandom(b.width)}
	}

	return pos, dir
}

func (b *Board) Init() {
	b.width = BOARD_WIDTH / BLOCK_SIZE
	b.height = BOARD_HEIGHT / BLOCK_SIZE

	b.grid = make([][]*Drawable, b.width)
	for r := 0; r < b.width; r++ {
		b.grid[r] = make([]*Drawable, b.height)
		for c := 0; c < b.height; c++ {

		}
	}
}

func (b *Board) GetData() []byte {
	var datas []byte
	for r := 0; r < b.width; r++ {
		for c := 0; c < b.height; c++ {
			if b.grid[r][c] != nil {
				obj := b.grid[r][c]
				data, err := json.Marshal(obj)
				if err != nil {
					log.Error("json marshal error: %v", err)
				}
				data = append(data, byte(','))
				datas = append(datas, data...)
			}
		}
	}
	return datas
}

func GetRandom(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
