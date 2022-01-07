package main

import (
	"github.com/glaukiol1/golssh/server/client/src"
)

func main() {
	src.Authenticate("PASSWORD", "localhost")
}
