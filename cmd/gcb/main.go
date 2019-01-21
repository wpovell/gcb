package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"gcb/bar"
	"gcb/blocks/bat"
	"gcb/blocks/music"
	"gcb/blocks/time"
	"gcb/blocks/wifi"
	"gcb/ipc"
	"gcb/log"
)

// Create bar, add blocks, and start
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	b := bar.New(ctx)

	// Create blocks
	timeBlk := time.New(b)
	wifiBlk := wifi.New(b)
	batBlk := bat.New(b)
	musicBlk := music.New(b)

	// Add blocks to bar
	b.AddBlock(bar.Center, timeBlk, "time")
	b.AddBlock(bar.Right, musicBlk, "music")
	b.AddBlock(bar.Right, wifiBlk, "wifi")
	b.AddBlock(bar.Right, batBlk, "bat")

	// Start IPC and bar
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go ipc.Start(b, ctx, wg)
	go b.Start(wg)

	// Wait for ctrl-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Cancel all goroutines
	log.Log("\nStopping", "stop")
	cancel()
	wg.Wait()
	log.Log("Done", "stop")
}
