package wrapper

import (
	"image"
	"time"

	"gcb/bar"
	"gcb/config"
	"gcb/text"

	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"golang.org/x/image/font"
)

type TextData struct {
	text  []string
	color []xgraphics.BGRA
}

func NewTextData() *TextData {
	return &TextData{
		text:  make([]string, 0),
		color: make([]xgraphics.BGRA, 0),
	}
}

func (t *TextData) Color(txt string, color xgraphics.BGRA) *TextData {
	t.text = append(t.text, txt)
	t.color = append(t.color, color)
	return t
}

func (t *TextData) Text(txt string) *TextData {
	t.text = append(t.text, txt)
	t.color = append(t.color, config.FG)
	return t
}

type TextBlock interface {
	Interval() time.Duration
	Text() *TextData
	Handle(ev xevent.ButtonPressEvent)
}

type TextW struct {
	bar *bar.Bar
	sub TextBlock
}

func NewTextW(b *bar.Bar, sub TextBlock) *TextW {
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

	for _, txt := range state.txt.text {
		state.width += font.MeasureString(state.drawer.Face, txt).Ceil()
	}
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
	txt    *TextData
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

	for i, text := range s.txt.text {
		color := s.txt.color[i]
		s.drawer.Src = image.NewUniform(color)
		s.drawer.DrawString(text)
	}
}
