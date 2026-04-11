package directory

import (
	"os"
)

func GetCurrent() string {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return p
}
