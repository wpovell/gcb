package ipc

import (
	"fmt"
	"net"
)

// Send `msg` to block called `name`
func Send(name, msg string) error {
	c, err := net.Dial("unix", sock_name)
	if err != nil {
		return err
	}

	_, err = c.Write([]byte(fmt.Sprintf("%s %s", name, msg)))
	if err != nil {
		return err
	}

	err = c.Close()
	return err
}
