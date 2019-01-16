package bat

import (
	"fmt"
	"time"

	"gcb/bar"
	"gcb/text"

	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"golang.org/x/image/font"
)

type Bat struct {
	bar    *bar.Bar
	drawer *font.Drawer
	txt    string
	ticker *time.Ticker
}

func Create(b *bar.Bar) *Bat {
	return &Bat{
		bar:    b,
		drawer: text.Drawer(),
	}
}

func (b *Bat) Handle(ev xevent.ButtonPressEvent) {}
func (b *Bat) Draw(x int, img *xgraphics.Image) {
	b.drawer.Dst = img
	b.drawer.Dot = text.Point(x)
	b.drawer.DrawString(b.txt)
}

func (b *Bat) Width() int {
	return font.MeasureString(b.drawer.Face, b.txt).Ceil()
}

func (b *Bat) Start() {
	b.ticker = time.NewTicker(time.Second)
	go func() {
		for {
			<-b.ticker.C
			bat := info()
			if bat.state == Charging {
				b.txt = fmt.Sprintf("Charging %d%%", bat.charge)
			} else {
				b.txt = fmt.Sprintf("Discharging %d%%", bat.charge)
			}
			b.bar.Redraw <- b
		}
	}()
}
