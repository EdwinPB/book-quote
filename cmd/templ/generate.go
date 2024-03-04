package templ

import (
	"os"
	"os/exec"
)

func Generate() {
	cmd := exec.Command("templ", "generate")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
