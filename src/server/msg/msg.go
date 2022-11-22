package msg

import (
	"server/base"

	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Register{})
	Processor.Register(&RegisterResp{})
	Processor.Register(&DataSync{})
	Processor.Register(&GameControl{})
	Processor.Register(&ChatMsg{})
}

type ChatMsg struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
	Date string `json:"date"`
}

type Register struct {
	Name string
}

type RegisterResp struct {
	Status bool      `json:"status"`
	Data   base.User `json:"data"`
}

type DataSync struct {
	DataMap string `json:"dataMap"`
}

type GameControl struct {
	Direction base.Direction `json:"direction"`
	Skill     base.GameKey   `json:"skill"`
}
