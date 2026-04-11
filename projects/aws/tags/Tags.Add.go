package tags

// Add - Add a new element to the map of key-value pairs
//
//	(c) 2023 Sam Caldwell.  MIT License
func (tags *Tags) Add(key, value string) {
	if tags.data == nil {
		tags.data = make(map[string]string)
	}
	//ToDo: validate the key
	tags.data[key] = value
}
