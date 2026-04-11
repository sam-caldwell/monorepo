package list

// Contains - Return whether value exists in any the sourceList array.
func Contains(sourceList []string, value string) bool {

    for _, item := range sourceList {

        if item == value {

            return true

        }

    }

    return false

}
