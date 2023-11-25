package pkg

import (
	"os/exec"
)

func ReformatFile(filename string) {
	c := exec.Command("gofmt", "-w", filename)
	if err := c.Run(); err != nil {
		panic(err)
	}
}
