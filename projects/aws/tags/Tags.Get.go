package tags

// Get - Given a key, return the associated value
//
//	(c) 2023 Sam Caldwell.  MIT License
func (tags *Tags) Get(key string) string {
	return tags.data[key]
}
