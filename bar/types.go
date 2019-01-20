package bar

import (
	"sync"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

const (
	LeftClick xproto.Button = iota + 1
	MiddleClick
	RightClick
	ScrollUp
	ScrollDown
)

// DrawState are passed to Bar to request a redraw
// They contain all needed data to draw
type DrawState interface {
	Draw(x int, img *xgraphics.Image)
	Width() int
	Source() Block
}

// Blocks handle events and send DrawState to bar as necessary
type Block interface {
	// Handle click
	Handle(ev xevent.ButtonPressEvent)
	// Start any update loop
	Start(wg *sync.WaitGroup) DrawState
}
