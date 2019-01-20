package bar

import (
	"gcb/config"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Return block associated with `x` coordinate
// Returns nil if no block associated
func (bar *Bar) findBlock(x int) Block {
	bar.Lock()
	defer bar.Unlock()
	for b, state := range bar.blocks {
		if state.Contains(x) {
			return b
		}
	}

	return nil
}

// Loop waiting for redraw requests
func (bar *Bar) Start() {
	go func() {
		// Start blocks
		for b, _ := range bar.blocks {
			state := &BlockState{
				state: b.Start(),
				start: 0,
			}
			bar.blocks[b] = state
		}

		// Handle redraw requests
		var state DrawState = nil
		bar.Lock()
		for {
			bar.draw()
			bar.Unlock()
			state = <-bar.Redraw
			bar.Lock()
			bar.blocks[state.Source()].state = state
		}
	}()
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
