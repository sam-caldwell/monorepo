package Disassember

type DisassemblerMethod int

const (
	linear DisassemblerMethod = iota
	recursive
)

const (
	strMethodLinear    = "linear"
	strMethodRecursive = "recursive"
)
