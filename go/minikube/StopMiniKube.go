package minikube

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
	"os/exec"
)

// Stop - stop the minikube service.
func Stop() {
	ansi.Cyan().Println("Stopping minikube").Reset()
	cmd := exec.Command("minikube", "delete", "--all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		ansi.Red().Printf("minikube stop failed.  error: %s", err).Fatal(1).Reset()
	}
	ansi.Green().Println("minikube stop ...success").Reset()
}
