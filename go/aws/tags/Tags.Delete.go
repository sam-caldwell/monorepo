package tags

// Delete - Delete the given map item
func (tags *Tags) Delete(key string) {
	delete(tags.data, key)
}
