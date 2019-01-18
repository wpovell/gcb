package main

import (
	"gcb/bar"
	"gcb/blocks/bat"
	"gcb/blocks/music"
	"gcb/blocks/time"
	"gcb/blocks/wifi"
)

// Create bar, add blocks, and start
func main() {
	b := bar.Create()

	timeBlk := time.Create(b)
	wifiBlk := wifi.Create(b)
	batBlk := bat.Create(b)
	musicBlk := music.Create(b)

	b.AddBlock(bar.Center, timeBlk)
	b.AddBlock(bar.Right, musicBlk)
	b.AddBlock(bar.Right, wifiBlk)
	b.AddBlock(bar.Right, batBlk)

	b.Start()

	select {}
}
