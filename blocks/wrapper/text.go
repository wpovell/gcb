package wrapper

import (
	"time"

	"gcb/bar"
	"gcb/text"

	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"golang.org/x/image/font"
)

type TextBlock interface {
	Interval() time.Duration
	Text() string
	Handle(ev xevent.ButtonPressEvent)
}

type TextW struct {
	bar    *bar.Bar
	drawer *font.Drawer
	sub    TextBlock
	txt    string
}

func CreateTextW(b *bar.Bar, sub TextBlock) *TextW {
	return &TextW{
		bar:    b,
		drawer: text.Drawer(),
		sub:    sub,
	}
}

func (t *TextW) Handle(ev xevent.ButtonPressEvent) {
	t.sub.Handle(ev)
}

func (t *TextW) Width() int {
	return font.MeasureString(t.drawer.Face, t.txt).Ceil()
}

func (t *TextW) Draw(x int, img *xgraphics.Image) {
	t.drawer.Dst = img
	t.drawer.Dot = text.Point(x)
	t.drawer.DrawString(t.txt)
}

func (t *TextW) Start() {
	go func() {
		for {
			t.txt = t.sub.Text()
			t.bar.Redraw <- t
			time.Sleep(t.sub.Interval())
		}
	}()
}
