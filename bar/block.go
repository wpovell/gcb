package bar

import (
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Interface for blocks
type Block interface {
	// Handle click
	Handle(ev xevent.ButtonPressEvent)
	// Draw itself
	Draw(x int, img *xgraphics.Image)
	// Return width (may be dynamic)
	Width() int
	// Start any update loop
	Start()
}
