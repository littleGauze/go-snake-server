package internal

import (
	"server/base"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	Game     = new(GameMain)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	go Game.Init()
}

func (m *Module) OnDestroy() {

}
