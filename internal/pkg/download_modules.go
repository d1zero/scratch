package pkg

import "os/exec"

func DownloadModules() {
	c := exec.Command("go", "mod", "tidy")
	if err := c.Run(); err != nil {
		panic(err)
	}
}
