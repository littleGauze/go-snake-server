package internal

import (
	"reflect"
	"server/base"
	"server/game"
	"server/msg"
	"strconv"
	"time"

	"github.com/name5566/leaf/gate"
)

func init() {
	handler(&msg.Register{}, handleAddUser)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleAddUser(args []interface{}) {
	m := args[0].(*msg.Register)
	a := args[1].(gate.Agent)

	user := newUser(m.Name)
	a.SetUserData(user)

	game.ChanRPC.Go("AddUser", a)

	a.WriteMsg(&msg.RegisterResp{Status: true, Data: user})
}

func newUser(name string) base.User {
	var user base.User
	user.Name = name
	user.Expires = 2 * 3600
	user.Token = strconv.Itoa(int(time.Now().UnixMicro()))
	return user
}
