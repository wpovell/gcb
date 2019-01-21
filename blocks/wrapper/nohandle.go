package wrapper

import (
	"gcb/bar"
)

type NoHandle struct {
	NoClick
	NoMsg
}

type NoClick struct{}

func (n *NoClick) HandleClick(ev bar.ClickEvent) bool { return false }

type NoMsg struct{}

func (n *NoMsg) HandleMsg(ev bar.MsgEvent) bool { return false }
