package main

import (
	"fmt"
	"os"
	"strings"

	"gcb/ipc"
)

func main() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Must supply a name and a message\n")
		os.Exit(1)
	}

	err := ipc.Send(os.Args[1], strings.Join(os.Args[2:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send message: %s", err.Error())
		os.Exit(1)
	}
}
