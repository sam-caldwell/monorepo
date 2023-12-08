package monorepo

import "time"

type Monorepo struct {
	StartTime    time.Time
	Root         string
	Debug        bool
	manifestList []Manifest
}
