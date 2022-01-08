package src

import (
	"bufio"
	"fmt"
	"net"

	"github.com/glaukiol1/golssh/server/client/helpers"
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
		helpers.Shell("127.0.0.1:9999")
	} else {
		println("Failed authentication")
	}
}
