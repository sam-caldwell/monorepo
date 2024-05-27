package configuration

// Init - Initialize the map
//
//	(c) 2023 Sam Caldwell.  MIT License.
func (cfg *Map[K, V]) Init() {
	*cfg = make(map[K]V)
}
