package configuration

// ExpectOrIgnore - Lookup a given key and return the value or default state if not present.
//
//	If the given key does not exist or the map is not initialized
//	the method will return the default value state
//
//	(c) 2023 Sam Caldwell.  MIT License
func (cfg *Map[K, V]) ExpectOrIgnore(name K) V {
	var defaultValue V
	if cfg != nil {
		if record, ok := (*cfg)[name]; ok {
			return record
		}
	}
	return defaultValue
}
