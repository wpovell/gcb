package main

import (
	"gcb/bar"
	"gcb/blocks/time"
)

// Create bar, add blocks, and start
func main() {
	b := bar.Create()
	timeBlk := time.Create(b)

	b.AddBlock(bar.Left, timeBlk)
	b.AddBlock(bar.Center, timeBlk)
	b.AddBlock(bar.Right, timeBlk)

	b.Start()
	timeBlk.Start()

	select {}
}
