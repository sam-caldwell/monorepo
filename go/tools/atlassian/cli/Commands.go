package cli

// Commands - a numeric representation of the command argument
type Commands uint8

const (
	Create Commands = iota
	Read
	Update
	Delete
	List
)
