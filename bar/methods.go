package bar

import (
	"gcb/config"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Return block associated with `x` coordinate
// Returns nil if no block associated
func (bar *Bar) findBlock(x int) Block {
	curX := 0
	for _, arr := range bar.blocks {
		for _, block := range arr {
			w := block.Width()
			if curX <= x && x <= curX+w {
				return block
			}
			curX += w
		}
	}

	return nil
}

// Loop waiting for redraw requests
func (bar *Bar) Start() {
	go func() {
		// Start blocks
		for _, arr := range bar.blocks {
			for _, block := range arr {
				block.Start()
			}
		}

		// Handle redraw requests
		for {
			<-bar.Redraw
			bar.draw()
		}
	}()
}

// Add block to position on bar
func (bar *Bar) AddBlock(pos Align, blk Block) {
	bar.blocks[pos] = append(bar.blocks[pos], blk)
}

// Draw entire bar
func (bar *Bar) draw() {
	// Background
	bar.img.For(func(x, y int) xgraphics.BGRA {
		return config.BG
	})

	var start int
	// Left
	start = 0
	for _, block := range bar.blocks[Left] {
		w := block.Width()
		block.Draw(start, bar.img)
		start += w
	}

	// Right
	l := len(bar.blocks[Right])
	start = bar.w
	for i := range bar.blocks[Right] {
		block := bar.blocks[Right][l-i-1]
		w := block.Width()
		start -= w
		block.Draw(start, bar.img)
	}

	// Center
	var total int
	for _, block := range bar.blocks[Center] {
		total += block.Width()
	}
	start = bar.w/2 - total/2
	for _, block := range bar.blocks[Center] {
		w := block.Width()
		block.Draw(start, bar.img)
		start += w
	}

	// Redraw
	bar.img.XDraw()
	bar.img.XPaint(bar.win.Id)
}
