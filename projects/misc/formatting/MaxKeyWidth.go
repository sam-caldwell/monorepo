package misc

func MaxKeyWidth(data map[string]string) (keyWidth int) {

	// Find the maximum key length
	for key := range data {
		if len(key) > keyWidth {
			keyWidth = len(key)
		}
	}
	return keyWidth
}
