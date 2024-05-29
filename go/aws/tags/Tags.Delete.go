package tags

// Delete - Delete the given map item
//
//	(c) 2023 Sam Caldwell.  MIT License
func (tags *Tags) Delete(key string) {
	delete(tags.data, key)
}
