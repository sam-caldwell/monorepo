//go:build darwin || windows

package minikube

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
	"os/exec"
)

// Start - Start the minikube service
func Start() {
	ansi.Cyan().Println("Starting minikube").Reset()
	cmd := exec.Command("minikube", "start", "--driver=docker")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		ansi.Red().Printf("minikube start failed.  error: %s", err).Fatal(1).Reset()
	}
	ansi.Green().Println("minikube start ...success").Reset()
}
