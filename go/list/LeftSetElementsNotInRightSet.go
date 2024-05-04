package list

import "fmt"

// LeftSetElementsNotInRightSet - Verify that all elements of *a are in *b
func LeftSetElementsNotInRightSet[T comparable](a *[]T, b *[]T) error {
	for _, elementA := range *a {
		found := false
		for _, elementB := range *b {
			if elementA == elementB {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid element '%v'", elementA)
		}
	}
	return nil
}
