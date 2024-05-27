package configuration

// Map - Provide a map-based configuration object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Map[Key comparable, Value any] map[Key]Value
