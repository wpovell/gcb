package ipc

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"gcb/bar"
	"gcb/log"
)

var ln net.Listener

func handle(cn net.Conn) {
	rdr := bufio.NewReader(cn)
	defer cn.Close()
	for {
		str, err := rdr.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Printf(str)
	}
}

func Start(bar *bar.Bar, ctx context.Context, wg *sync.WaitGroup) {
	var err error
	os.Remove("/tmp/gcb.sock")
	ln, err = net.Listen("unix", "/tmp/gcb.sock")
	log.Fatal(err)

	// Cancellation
	go func() {
		<-ctx.Done()
		log.Log("Stopped IPC", "stop")
		ln.Close()
		wg.Done()
	}()

	// Accept loop
	go func() {
		for {
			cn, err := ln.Accept()
			if err != nil {
				return
			}
			go handle(cn)
		}
	}()
}
