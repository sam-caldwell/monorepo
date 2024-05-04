package list

import "fmt"

// LeftSetElementsNotInRightSet - Verify that all elements of *a are in *b
func LeftSetElementsNotInRightSet(a *[]any, b *[]any) error {
	for _, elementA := range *a {
		found := false
		for _, elementB := range *b {
			if elementA == elementB {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid element '%s'", elementA)
		}
	}
	return nil
}
