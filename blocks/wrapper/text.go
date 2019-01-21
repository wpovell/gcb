package wrapper

import (
	"fmt"
	"image"
	"sync"
	"time"

	"gcb/bar"
	"gcb/config"
	"gcb/ipc"
	"gcb/log"
	"gcb/text"

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
	HandleClick(e bar.ClickEvent) bool
	HandleMsg(m ipc.MsgEvent) bool
}

type TextW struct {
	bar    *bar.Bar
	sub    TextBlock
	events chan interface{}
}

func NewTextW(b *bar.Bar, sub TextBlock) *TextW {
	return &TextW{
		bar:    b,
		sub:    sub,
		events: make(chan interface{}),
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

func (t *TextW) EventCh() chan interface{} {
	return t.events
}

func (t *TextW) Start(wg *sync.WaitGroup) bar.DrawState {
	wg.Add(1)
	go func() {
		timer := time.NewTimer(t.sub.Interval())
		for {
			select {
			case ev := <-t.events:
				if t.handle(ev) {
					t.bar.Redraw <- t.createState()
				}
			case <-timer.C:
				t.bar.Redraw <- t.createState()
				timer = time.NewTimer(t.sub.Interval())
			case <-t.bar.Ctx.Done():
				log.Log(fmt.Sprintf("Stopping %T\n", t.sub))
				wg.Done()
				return
			}
		}
	}()
	return t.createState()
}

func (t *TextW) handle(ev interface{}) bool {
	switch ev.(type) {
	case bar.ClickEvent:
		return t.sub.HandleClick(ev.(bar.ClickEvent))
	case ipc.MsgEvent:
		return t.sub.HandleMsg(ev.(ipc.MsgEvent))
	}
	panic("Bad event type")
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
