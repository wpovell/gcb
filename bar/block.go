package bar

import (
	"github.com/BurntSushi/xgbutil/xevent"
  "github.com/BurntSushi/xgbutil/xgraphics"
)

type Block interface {
	Handle(ev xevent.ButtonPressEvent)
	Draw(x int, img *xgraphics.Image)
	Width() int
	Start()
}
