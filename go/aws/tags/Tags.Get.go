package tags

// Get - Given a key, return the associated value
func (tags *Tags) Get(key string) string {
	return tags.data[key]
}
