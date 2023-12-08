package pkg

import (
	"fmt"
	"os/exec"
)

func DownloadModules(serviceName string) {
	c := exec.Command(fmt.Sprintf("cd %s && go mod tidy", serviceName))
	if err := c.Run(); err != nil {
		panic(err)
	}
}
