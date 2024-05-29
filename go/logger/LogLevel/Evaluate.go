package LogLevel

// Evaluate - Evaluate internal log state against expectedLevel
//
//	(c) 2023 Sam Caldwell.  MIT License
func (level *Value) Evaluate(expectedLevel Value) bool {
	if expectedLevel <= *level {
		return true
	}
	return false
}
