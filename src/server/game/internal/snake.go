package internal

import (
	"server/base"

	"github.com/name5566/leaf/gate"
)

const (
	DEFAULT_LENGTH int = 3
	JUMP_DISTANCE  int = 3
)

type Snake struct {
	Agent *gate.Agent

	SnakeSegment
	skipNextTurn bool
	hitDetected  bool
	IsAlive      bool `json:"isAlive"`

	Speed     Speed          `json:"speed"`
	Direction base.Direction `json:"direction"`

	hiscore int
	Points  int `json:"points"`
	lives   int

	Name      string `json:"name"`
	Token     string `json:"token"`
	segments  []Drawable
	MaxLength int `json:"maxLength"`
}

func (s *Snake) Init(pos Position, dir base.Direction, user base.User) {
	s.SnakeSegment = SnakeSegment{}

	s.MaxLength = DEFAULT_LENGTH
	s.segments = make([]Drawable, 1)
	s.segments[0] = s
	s.IsAlive = true
	s.Clazz = "Snake"
	s.Name = user.Name
	s.Token = user.Token
	s.lives = 999
	s.Direction = dir

	Game.Board.PlaceObject(s, pos)
}

func (s *Snake) Die() {
	s.hitDetected = true
	if s.Points > s.hiscore {
		s.hiscore = s.Points
	}

	s.lives -= 1
	s.destroy()

	// s.Position = Position{X: 0, Y: 0}
	s.Direction = base.NONE
}

func (s *Snake) destroy() {
	for _, seg := range s.segments {
		Game.Board.RemoveObject(seg.GetPosition())
	}

	s.segments = make([]Drawable, 1)
	s.segments[0] = s
	s.MaxLength = DEFAULT_LENGTH
	s.Points = 0

	s.Remove()
}

func (s *Snake) SetSpeed(speed Speed) {
	s.Speed = speed
	s.skipNextTurn = (speed == SLOW)
}

func (s *Snake) ProcessTurn() {
	if !s.IsAlive {
		return
	}

	if s.Speed != FAST && Game.Clock.Tick == TICK {
		return
	}

	if s.Speed == SLOW && Game.Clock.Tick == TOCK {
		s.skipNextTurn = !s.skipNextTurn
		if s.skipNextTurn {
			return
		}
	}

	s.hitDetected = false

	isMoving := true
	pos := s.Position.ClonePosition()

	switch s.Direction {
	case base.UP:
		pos.Y -= 1
	case base.DOWN:
		pos.Y += 1
	case base.LEFT:
		pos.X -= 1
	case base.RIGHT:
		pos.X += 1
	default:
		isMoving = false
	}

	if isMoving {
		if pos.X < 0 {
			pos.X = Game.Board.width - 1
		} else if pos.X == Game.Board.width {
			pos.X = 0
		} else if pos.Y == Game.Board.height {
			pos.Y = 0
		} else if pos.Y < 0 {
			pos.Y = Game.Board.height - 1
		}

		if Game.Board.grid[pos.X][pos.Y] != nil {
			obj := Game.Board.grid[pos.X][pos.Y]
			(*obj).HandleCollision(s)
		}

		if !s.IsAlive {
			s.destroy()
		} else {
			s.update(pos)
		}
	}
}

func (s *Snake) update(pos Position) {
	lastPos := s.GetPosition().ClonePosition()
	for i, seg := range s.segments {
		var newPos Position
		if i == 0 {
			newPos = pos
		} else {
			newPos = lastPos
		}
		lastPos = seg.GetPosition().ClonePosition()
		Game.Board.MoveObject(seg, newPos)
	}

	if len(s.segments) < s.MaxLength {
		seg := SnakeSegment{}
		seg.Init(lastPos)
		s.segments = append(s.segments, &seg)
		Game.Board.PlaceObject(&seg, lastPos)
	}
}

func (s *Snake) Remove() {
	Game.Board.RemoveObject(s.Position)
}

func (s *Snake) SetDirection(dir base.Direction) {
	s.Direction = dir
}
