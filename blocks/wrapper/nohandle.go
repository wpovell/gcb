package wrapper

import (
	"gcb/bar"
	"gcb/ipc"
)

// Ignore all events
type NoHandle struct {
	NoClick
	NoMsg
}

// Ignore click events
type NoClick struct{}

func (n *NoClick) HandleClick(ev bar.ClickEvent) bool { return false }

// Ignore message events
type NoMsg struct{}

func (n *NoMsg) HandleMsg(ev ipc.MsgEvent) bool { return false }
