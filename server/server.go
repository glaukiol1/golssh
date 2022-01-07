package main

import (
	"fmt"
	"net"

	tcpserver "github.com/glaukiol1/golssh/server/server/tcpserver"
)

func main() {
	PORT := ":" + fmt.Sprint(22)
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go tcpserver.HandleConnection(c)
	}
}
