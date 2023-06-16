package misc

func CalcMaxColumnWidth(list *[]string) (maxWidth int) {
	if list != nil {
		for _, line := range *list {
			if width := len(line); width > maxWidth {
				maxWidth = len(line)
			}
		}
	}
	return maxWidth
}
