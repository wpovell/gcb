package ipc

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"gcb/bar"
	"gcb/log"
)

type MsgEvent struct {
	Name string
	Msg  string
}

const sock_name = "/tmp/gcb.sock"

// Handle a socket connection
func handle(cn net.Conn, bar *bar.Bar) {
	rdr := bufio.NewReader(cn)
	defer cn.Close()
	for {
		str, err := rdr.ReadString('\n')
		if err != nil {
			return
		}

		parts := strings.SplitN(strings.TrimSpace(str), " ", 2)
		name := parts[0]

		if len(parts) != 2 || bar.Names[name] == nil {
			log.Log(fmt.Sprintf("Invalid message: %#v", parts), "warn", "ipc")
			return
		}
		msg := parts[1]
		bar.Names[name].EventCh() <- MsgEvent{name, msg}
	}
}

// Start IPC handling
func Start(bar *bar.Bar, ctx context.Context, wg *sync.WaitGroup) {
	var err error
	// Purge old socket if it failed to be removed
	os.Remove(sock_name)
	ln, err := net.Listen("unix", sock_name)
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
			go handle(cn, bar)
		}
	}()
}
