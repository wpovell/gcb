package bar

import (
	"sync"

	"gcb/config"
	"gcb/log"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Return block associated with `x` coordinate
// Returns nil if no block associated
func (bar *Bar) findBlock(x int) Block {
	for b, state := range bar.blocks {
		if state.Contains(x) {
			return b
		}
	}

	return nil
}

// Loop waiting for redraw requests
func (bar *Bar) Start(wg *sync.WaitGroup) {
	// Start blocks
	for b, _ := range bar.blocks {
		state := &BlockState{
			state: b.Start(wg),
			start: 0,
		}
		bar.blocks[b] = state
	}

	// Handle redraw requests
	bar.draw()
	for {
		select {
		case ev := <-bar.xevent:
			block := bar.findBlock(int(ev.EventX))
			block.EventCh() <- ClickEvent{
				Button: ev.Detail,
				X:      int(ev.EventX),
				Y:      int(ev.EventY),
			}
		case state := <-bar.Redraw:
			bar.blocks[state.Source()].state = state
			bar.draw()
		case <-bar.Ctx.Done():
			log.Log("Stopping bar", "stop")
			wg.Done()
			return
		}
	}
}

// Add block to position on bar
func (bar *Bar) AddBlock(aln Align, blk Block) {
	bar.blocks[blk] = nil
	bar.align[aln] = append(bar.align[aln], blk)
}

// Draw entire bar
func (bar *Bar) draw() {
	// Background
	bar.img.For(func(x, y int) xgraphics.BGRA {
		return config.BG
	})

	var start int
	// Left
	posn := config.Padding
	for _, block := range bar.align[Left] {
		bs := bar.blocks[block]
		bs.Draw(posn, bar.img)
		start += bs.state.Width()
		start += config.Padding
	}

	// Right
	l := len(bar.align[Right])
	posn = bar.w
	for i := range bar.align[Right] {
		block := bar.align[Right][l-i-1]
		bs := bar.blocks[block]
		posn -= bs.state.Width()
		posn -= config.Padding
		bs.Draw(posn, bar.img)
	}

	// Center
	var total int
	for _, block := range bar.align[Center] {
		bs := bar.blocks[block]
		total += bs.state.Width()
		total += config.Padding
	}
	start = bar.w/2 - total/2
	for _, block := range bar.align[Center] {
		bs := bar.blocks[block]
		bs.Draw(start, bar.img)
		start += bs.state.Width()
		start += config.Padding
	}

	// Redraw
	bar.img.XDraw()
	bar.img.XPaint(bar.win.Id)
}
