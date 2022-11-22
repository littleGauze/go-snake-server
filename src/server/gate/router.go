package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Register{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.GameControl{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ChatMsg{}, game.ChanRPC)
}
