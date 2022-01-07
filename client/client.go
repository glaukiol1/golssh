package main

import (
	"github.com/glaukiol1/golssh/server/client/src"
)

func main() {
	src.Connect("127.0.0.1:22", "PASSWORD")
}
