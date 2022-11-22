package internal

import (
	"reflect"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("AddUser", handleAddUser)
	handler(&msg.GameControl{}, handleGameControl)
	handler(&msg.ChatMsg{}, handleChatMsg)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func handleAddUser(args []interface{}) {
	a := args[0].(gate.Agent)
	Game.addPlayer(a)
}
func handleGameControl(args []interface{}) {
	m := args[0].(*msg.GameControl)
	a := args[1].(gate.Agent)

	Game.SetDirection(a, m.Direction)
}

func handleChatMsg(args []interface{}) {
	m := args[0].(*msg.ChatMsg)

	for _, p := range Game.players {
		(*p.Agent).WriteMsg(&msg.ChatMsg{Msg: m.Msg, Name: p.Name, Date: time.Now().Local().Format("2006-01-02 15:04:05")})
	}
}
