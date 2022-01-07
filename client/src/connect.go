package src

import (
	"bufio"
	"fmt"
	"net"
)

func Connect(host, password string) {
	CONNECT := host
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Write([]byte(password + "\n"))
	message, _ := bufio.NewReader(c).ReadString('\n')
	if message == "success" {
		println("Successful authentication!")
	} else {
		println("Failed authentication")
	}
}
