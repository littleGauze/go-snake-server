package internal

import (
	"time"
)

type ClockType int

const (
	TIMED = iota
	INFINITE
)

type ClockTick int

const (
	TICK = iota
	TOCK
)

type Handler = func()

type Timer struct {
	interval int64
	timeType ClockType
	isRuning bool
	isPaused bool

	stop    chan bool
	timer   *time.Timer
	ticker  *time.Ticker
	handler Handler

	Tick ClockTick
}

func (t *Timer) NewTimer(interval int64, isInfinite bool, handler Handler) {
	var tt ClockType = TIMED
	if isInfinite {
		tt = INFINITE
	}
	t.interval = interval
	t.handler = handler
	t.timeType = tt
	t.stop = make(chan bool)
	t.Tick = TICK
}

func (t *Timer) Pause()  { t.isPaused = true }
func (t *Timer) Resume() { t.isPaused = false }

func (t *Timer) Start() {
	t.isRuning = true
	if t.timeType == TIMED {
		t.timer = time.AfterFunc(time.Duration(t.interval)*time.Millisecond, t.onElapsed)
	} else {
		go t.handleInterval(time.Duration(t.interval) * time.Millisecond)
	}
}

func (t *Timer) Stop() {
	t.isRuning = false
	if t.timeType == TIMED {
		t.timer.Stop()
	} else {
		t.ticker.Stop()
		t.stop <- true
	}
}

func (t *Timer) handleInterval(d time.Duration) {
	t.ticker = time.NewTicker(d)

	for {
		select {
		case <-t.stop:
			return
		case <-t.ticker.C:
			t.onElapsed()
		}
	}
}

func (t *Timer) onElapsed() {
	if t.isPaused {
		return
	}

	if t.Tick == TICK {
		t.Tick = TOCK
	} else {
		t.Tick = TICK
	}

	t.handler()

	if t.timeType == TIMED {
		t.timer.Stop()
	}
}
