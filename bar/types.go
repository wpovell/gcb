package bar

import (
	"sync"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Block alignment on bar
type Align int

const (
	Left   Align = 0
	Center Align = 1
	Right  Align = 2
)

// Mouse button enum
const (
	LeftClick xproto.Button = iota + 1
	MiddleClick
	RightClick
	ScrollUp
	ScrollDown
)

// Event passed to block on click
type ClickEvent struct {
	Button xproto.Button
	X, Y   int
}

// Current data and placement of a displayed block
type BlockState struct {
	state DrawState
	start int
}

func (s *BlockState) Contains(x int) bool {
	return s.start < x && x < s.start+s.state.Width()
}

func (s *BlockState) Draw(x int, img *xgraphics.Image) {
	s.start = x
	s.state.Draw(x, img)
}

// DrawState are passed to Bar to request a redraw
// They contain all needed data to draw
type DrawState interface {
	Draw(x int, img *xgraphics.Image)
	Width() int
	Source() Block
}

// Blocks handle events and send DrawState to bar as necessary
type Block interface {
	// Handle events
	EventCh() chan interface{}
	// Start any update loop
	Start(wg *sync.WaitGroup) DrawState
}
