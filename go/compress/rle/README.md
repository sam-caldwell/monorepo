Run-Length Encoding
===================

## Objective
Create [`Encode()`](./Encode.go) and [`Decode()`](./Decode.go) functions
which perform Run-Length encoding, a simple form of compression.


## Usage

GIVEN a `[]byte` slice input:
```go
input:=[]byte("AABBBCCCC")
```
WHERE we call `Encode(input)`

WE EXPECT the function to return:
```go
output:=[]byte("A2B3C4")
```
WHERE each repeated symbol is replaced by a single instance
of the symbol (`byte`) followed by a number of times the symbol is
repeated.

Likewise, `Decode(output)` will return the original `input` slice.

