package convert

// CapitalizeStringList - Capitalize the first letter in every element of a string array
//
//	(c) 2023 Sam Caldwell.  MIT License
func CapitalizeStringList(list *[]string) {
	for i := 0; i < len(*list); i++ {
		(*list)[i] = Capitalize((*list)[i])
	}
}
