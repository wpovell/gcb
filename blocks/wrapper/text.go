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
	bar *bar.Bar
	sub TextBlock
}

func CreateTextW(b *bar.Bar, sub TextBlock) *TextW {
	return &TextW{
		bar: b,
		sub: sub,
	}
}

func (t *TextW) createState() *TextWState {
	state := &TextWState{
		txt:    t.sub.Text(),
		block:  t,
		drawer: text.Drawer(),
	}
	state.width = font.MeasureString(state.drawer.Face, state.txt).Ceil()
	return state
}

func (t *TextW) Start() bar.DrawState {
	go func() {
		for {
			time.Sleep(t.sub.Interval())
			t.bar.Redraw <- t.createState()
		}
	}()
	return t.createState()
}

func (t *TextW) Handle(ev xevent.ButtonPressEvent) {
	t.sub.Handle(ev)
}

type TextWState struct {
	txt    string
	width  int
	block  *TextW
	drawer *font.Drawer
}

func (s *TextWState) Source() bar.Block {
	return s.block
}

func (s *TextWState) Width() int {
	return s.width
}

func (s *TextWState) Draw(x int, img *xgraphics.Image) {
	s.drawer.Dst = img
	s.drawer.Dot = text.Point(x)
	s.drawer.DrawString(s.txt)
}
