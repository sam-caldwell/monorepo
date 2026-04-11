package letters

import "strings"

type Letters string

func (L *Letters) String() string { return string(*L) }
func (L *Letters) Rune() rune     { return []rune((*L))[0] }
func (L *Letters) Runes() []rune  { return []rune((*L)) }
func (L *Letters) Lower()         { *L = Letters(strings.ToLower(string(*L))) }
func (L *Letters) Upper()         { *L = Letters(strings.ToUpper(string(*L))) }

const (
	A = "A"
	B = "B"
	C = "C"
	D = "D"
	E = "E"
	F = "F"
	G = "G"
	H = "H"
	I = "I"
	J = "J"
	K = "K"
	L = "L"
	M = "M"
	N = "N"
	O = "O"
	P = "P"
	Q = "Q"
	R = "R"
	S = "S"
	T = "T"
	U = "U"
	V = "V"
	W = "W"
	X = "X"
	Y = "Y"
	Z = "Z"
)
