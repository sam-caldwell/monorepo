package LogLevel

// Evaluate - Evaluate internal log state against expectedLevel
func (level *Value) Evaluate(expectedLevel Value) bool {
	if expectedLevel <= *level {
		return true
	}
	return false
}
