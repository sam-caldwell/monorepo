package configuration

import (
	"sync"
)

// Map - Provide a map-based configuration object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Map[Key comparable, Value any] struct {
	lock sync.RWMutex
	data map[Key]Value
}
