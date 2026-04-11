package types

// OperatorId - The numeric identifier for an operator used in a ThreatQlQuery
type OperatorId uint8 //e.g. 0=equals, 1=notEquals, ... n=regex

const (
	Equals OperatorId = iota
	NotEquals
	LessThan
	GreaterThan
	LessThanEquals
	GreaterThanEquals
	Regex
)
