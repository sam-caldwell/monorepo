package misc

func MaxKeyValueWidth(data map[string]string) (keyWidth, valueWidth int) {

	// Find the maximum key length
	for key, value := range data {
		if len(key) > keyWidth {
			keyWidth = len(key)
		}
		if len(value) > valueWidth {
			valueWidth = len(value)
		}
	}
	return keyWidth, valueWidth
}
