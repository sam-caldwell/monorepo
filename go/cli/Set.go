package cli

type Set map[string]bool

// Add - Add element to the set
func (s *Set) Add(e string) {
	if s == nil {
		*s = make(map[string]bool)
	}
	(*s)[e] = true
}

// ToArray - Convert the map to []string
func (s *Set) ToArray() (result []string) {
	// Convert the map to []string
	for i, _ := range *s {
		result = append(result, i)
	}
	return result
}
