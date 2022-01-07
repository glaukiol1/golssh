package tcpserver

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/glaukiol1/golssh/server/server/helpers"
)

func HandleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		println(temp)
		if temp == helpers.Password {
			c.Write([]byte(string("success")))
			print(strings.Split(c.RemoteAddr().String(), ":")[0])
			break
		} else {
			c.Write([]byte(string("error")))
			break
		}
	}
	c.Close()
}
