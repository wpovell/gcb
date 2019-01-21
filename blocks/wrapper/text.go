package wrapper

import (
	"fmt"
	"sync"
	"time"

	"gcb/bar"
	"gcb/ipc"
	"gcb/log"
	"gcb/text"

	"golang.org/x/image/font"
)

// Interface implemented by wrapped blocks
type TextBlock interface {
	Interval() time.Duration
	Text() *TextData
	HandleClick(e bar.ClickEvent) bool
	HandleMsg(m ipc.MsgEvent) bool
}

// Block wrapper
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

// Create a object for current state to pass to bar
func (t *TextW) createState() *TextWState {
	state := &TextWState{
		txt:    t.sub.Text(),
		block:  t,
		drawer: text.Drawer(),
	}

	// Calculate width
	for _, txt := range state.txt.text {
		state.width += font.MeasureString(state.drawer.Face, txt).Ceil()
	}
	return state
}

// Return channel which bar/ipc will pass events to
func (t *TextW) EventCh() chan interface{} {
	return t.events
}

// Block update loop
func (t *TextW) Start(wg *sync.WaitGroup) bar.DrawState {
	wg.Add(1)
	go func() {
		timer := time.NewTimer(t.sub.Interval())
		for {
			select {
			case ev := <-t.events:
				// Event
				if t.handle(ev) {
					t.bar.Redraw <- t.createState()
				}
			case <-timer.C:
				// Update
				t.bar.Redraw <- t.createState()
				timer = time.NewTimer(t.sub.Interval())
			case <-t.bar.Ctx.Done():
				// Done
				log.Log(fmt.Sprintf("Stopping %T\n", t.sub))
				wg.Done()
				return
			}
		}
	}()
	return t.createState()
}

// Handle an event recieved, passing to wrapped struct
func (t *TextW) handle(ev interface{}) bool {
	switch ev.(type) {
	case bar.ClickEvent:
		return t.sub.HandleClick(ev.(bar.ClickEvent))
	case ipc.MsgEvent:
		return t.sub.HandleMsg(ev.(ipc.MsgEvent))
	}
	panic("Bad event type")
}
