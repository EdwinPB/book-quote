package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("templ", "generate")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
