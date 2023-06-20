//go:build linux
// +build linux

package systemrecon

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetFamily() (output string, err error) {
	var file *os.File
	if file, err = os.Open("/etc/os-release"); err != nil {
		return output, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			distribution := strings.TrimPrefix(line, "ID=")
			return distribution, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return output, err
	}

	return output, errors.New("unable to detect Linux distribution")
}
