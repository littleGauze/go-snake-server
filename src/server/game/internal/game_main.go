package internal

import (
	"server/base"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

type GameDifficuty int

type DataCallback func(data string)

const (
	EASY     GameDifficuty = 200
	MEDIUM   GameDifficuty = 150
	DIFFICUT GameDifficuty = 50
)

type GameMain struct {
	Clock *Timer
	Board *Board

	players []*Snake

	HiScore int

	isReady     bool
	isRunning   bool
	coinCounter int
}

func (gm *GameMain) Init() {

	gm.Clock = &Timer{}
	gm.Board = &Board{}

	gm.Ready()
}

func (gm *GameMain) Ready() {
	// init board
	gm.Board.Init()

	gm.Clock.NewTimer(int64(DIFFICUT), true, func() {
		gm.onClockTick()

	})
}

func (gm *GameMain) Start() {
	if gm.isRunning {
		return
	}
	if gm.Clock.isPaused {
		gm.pause()
		return
	}

	gm.isRunning = true
	gm.Clock.Start()
}

func (gm *GameMain) pause() {
	if gm.Clock.isPaused {
		gm.isRunning = true
		gm.Clock.Resume()
		return
	}

	gm.isRunning = false
	gm.Clock.Pause()
}

func (gm *GameMain) reset() {
	gm.Clock.Stop()
	gm.isRunning = false
	gm.Ready()
}

func (gm *GameMain) onClockTick() {
	if gm.Clock.Tick == TOCK {
		gm.coinCounter++
		if gm.coinCounter >= 2 {
			gm.coinCounter = 0
		}

		if GetRandom(10) < 5 {
			p := (float64(CoinActive) + 0.5) / 15
			if (float64(GetRandom(10))*0.1 + p) < 0.5 {
				if GetRandom(10) < 6 {
					coin := CreateRandomCoin()
					gm.Board.PlaceAtRandom(&coin)
				} else if GetRandom(10) < 5 {
					coin := CreateRandomSpeedCoin(FAST)
					gm.Board.PlaceAtRandom(&coin)
				} else {
					coin := CreateRandomSpeedCoin(SLOW)
					gm.Board.PlaceAtRandom(&coin)
				}
			}
		}
	}

	data := gm.Board.GetData()
	for _, p := range gm.players {
		p.ProcessTurn()
		(*p.Agent).WriteMsg(&msg.DataSync{DataMap: string(data)})
	}
}

func (gm *GameMain) addPlayer(agent gate.Agent) {
	user := agent.UserData().(base.User)
	for _, p := range gm.players {
		if (*p.Agent).RemoteAddr() == agent.RemoteAddr() {
			return
		}
	}
	pos, dir := gm.Board.GetRandomPositionAndDirection()
	player := new(Snake)
	player.Init(pos, dir, user)
	player.Agent = &agent
	gm.players = append(gm.players, player)

	if !gm.isRunning {
		gm.Start()
	}
}

func (gm *GameMain) removePlayer(agent gate.Agent) {
	for i, p := range gm.players {
		if (*p.Agent).RemoteAddr() == agent.RemoteAddr() {
			gm.players = append(gm.players[:i], gm.players[i+1:]...)
			p.destroy()
		}
	}
}

func (gm *GameMain) SetDirection(agent gate.Agent, dir base.Direction) {
	for _, p := range gm.players {
		if (*p.Agent).RemoteAddr() == agent.RemoteAddr() {
			p.SetDirection(dir)
		}
	}
}
