package numbers

import (
	"strconv"
)

type Number int

func (n *Number) String() string { return strconv.Itoa(int(*n)) }
func (n *Number) Byte() byte     { return byte(*n) }
func (n *Number) Int() int       { return int(*n) }
func (n *Number) Int8() int8     { return int8(*n) }
func (n *Number) Int16() int16   { return int16(*n) }
func (n *Number) Int32() int32   { return int32(*n) }
func (n *Number) Int64() int64   { return int64(*n) }
func (n *Number) Uint() uint     { return uint(*n) }
func (n *Number) Uint8() uint8   { return uint8(*n) }
func (n *Number) Uint16() uint16 { return uint16(*n) }
func (n *Number) Uint32() uint32 { return uint32(*n) }
func (n *Number) Uint64() uint64 { return uint64(*n) }

func (n *Number) Add(d Number) *Number  { *n += d; return n }
func (n *Number) Dec() *Number          { *n++; return n }
func (n *Number) Div(d Number) *Number  { *n = *n / d; return n }
func (n *Number) Inc() *Number          { *n++; return n }
func (n *Number) Mult(d Number) *Number { *n *= d; return n }
func (n *Number) Sqr() *Number          { *n *= *n; return n }

const (
	Zero  Number = 0
	One   Number = 1
	Two   Number = 2
	Three Number = 3
	Four  Number = 4
	Five  Number = 5
	Six   Number = 6
	Seven Number = 7
	Eight Number = 8
	Nine  Number = 9
)

const (
	Float32zero = float32(0.0)
	Float64zero = 0.0
	IntZero     = 0
)
