package wifi

import (
	"time"

	"gcb/bar"
	"gcb/text"

	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"golang.org/x/image/font"
)

const (
	intf = "wlp4s0"
)

type Wifi struct {
	bar    *bar.Bar
	ticker *time.Ticker
	drawer *font.Drawer
	txt    string
}

func Create(b *bar.Bar) *Wifi {
	return &Wifi{
		bar:    b,
		ticker: nil,
		drawer: text.Drawer(),
	}
}

func (w *Wifi) Handle(ev xevent.ButtonPressEvent) {}
func (w *Wifi) Draw(x int, img *xgraphics.Image) {
	w.drawer.Dst = img
	w.drawer.Dot = text.Point(x)
	w.drawer.DrawString(w.txt)
}

func (w *Wifi) Width() int {
	return font.MeasureString(w.drawer.Face, w.txt).Ceil()
}

func (w *Wifi) Start() {
	w.ticker = time.NewTicker(time.Second)
	go func() {
		for {
			txt, err := ssid(intf)
			if err != nil {
				w.txt = "No Wifi"
			} else {
				w.txt = txt
			}
			w.bar.Redraw <- w
			<-w.ticker.C
		}
	}()
}
