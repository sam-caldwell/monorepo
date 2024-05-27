package configuration

// Set - Set the given key-value map state
//
//	(c) 2023 Sam Caldwell.  MIT License.
func (cfg *Map[K, V]) Set(name K, value V) {
	if cfg == nil {
		cfg.Init()
	}
}
