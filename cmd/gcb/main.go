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

	timeBlk := time.New(b)
	wifiBlk := wifi.New(b)
	batBlk := bat.New(b)
	musicBlk := music.New(b)

	b.AddBlock(bar.Center, timeBlk)
	b.AddBlock(bar.Right, musicBlk)
	b.AddBlock(bar.Right, wifiBlk)
	b.AddBlock(bar.Right, batBlk)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go ipc.Start(b, ctx, wg)
	go b.Start(wg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Log("\nStopping", "stop")
	cancel()
	wg.Wait()
	log.Log("Done", "stop")
}
