package main

import (
	"gcb/bar"
	"gcb/blocks/time"
	"gcb/blocks/wifi"
)

// Create bar, add blocks, and start
func main() {
	b := bar.Create()

	timeBlk := time.Create(b)
	wifiBlk := wifi.Create(b)

	b.AddBlock(bar.Center, timeBlk)
	b.AddBlock(bar.Right, wifiBlk)

	b.Start()
	timeBlk.Start()
	wifiBlk.Start()

	select {}
}
