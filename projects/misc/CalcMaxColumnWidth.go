package misc

/*
 * CalcMaxColumnWidth
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file creates a function which will take a list of
 * strings and calculate the width of the longest string,
 * returning the same as an int.
 *
 * See README.md
 */

// CalcMaxColumnWidth - given a list calculate the maximum width of the strings in the list.
func CalcMaxColumnWidth(list *[]string) (maxWidth int) {
	if list != nil {
		for _, line := range *list {
			if width := len(line); width > maxWidth {
				maxWidth = width
			}
		}
	}
	return maxWidth
}
