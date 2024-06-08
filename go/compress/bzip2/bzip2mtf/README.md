Move-to-Front (MTF) Transformation
==================================

## Objective

The Move-to-Front (MTF) transformation is a simple but powerful algorithm used primarily
in data compression. It transforms a sequence of input symbols into a sequence of indices,
which often results in a sequence that is easier to compress using other techniques,
such as [Run-Length Encoding (RLE)](../../rle/README.md) or [Huffman Coding](../../huffman/README.md).

## Explanation

### How MTF Works

#### Initialization:

> Start with a list of all possible symbols in the input alphabet, typically sorted in
> lexicographical order. For byte data, this list contains all 256 possible byte
> values (0 to 255).

#### Transformation:

> For each symbol in the input sequence, find its position (index) in the list.
> Output this index.
> Move the symbol to the front of the list.

## Applications

### Compression Algorithms

> MTF is often used as a preprocessing step in compression algorithms like the
> [Burrows-Wheeler Transform (BWT)](../bzip2bwt/Bwt.go) to improve their efficiency.
> By transforming the input into indices that often cluster around small values,
> MTF enhances the efficiency of subsequent compression stages, making it a valuable
> tool in the data compression toolkit.

### Data Streams

> It is also used in situations where adaptive coding techniques are applied to data streams.

## Benefits of MTF

* ***Locality***: MTF tends to output small indices for symbols that occur frequently or recently. This can make
  the transformed sequence more amenable to compression techniques that benefit from such patterns.


* ***Simplicity***: The algorithm is straightforward to implement and understand.

### Example

The following table shows the MTF process, given an input `"abacabad"`:

1. We start by initializing the `commonDictionary = [0, 1, 2, ..., 255]`
2. For each character (c) in the input string (`"abacabad"`), we perform the
   following, where position (p) is the position in the dictionary:

| c         | Dictionary                                              | Output                          | 
|-----------|---------------------------------------------------------|---------------------------------|
| 'a' (97)  | `[0, 1, 2, ..., 96, 98, ..., 255]`                      | `[97]`                          | 
| 'b' (98)  | `[0, 97, 1, 2, ..., 96, 98, 99, ..., 255]`              | `[97, 98]`                      |
| 'a' (97)  | `[97, 98, 1]`                                           | `[97, 98, 1]`                   |
| 'c' (99)  | `[0, 97, 1, 2, ..., 96, 98, 1, 100, ..., 255]`          | `[97, 98, 1, 99]`               |
| 'a' (97)  | `[0, 99, 97, 1, 2, ..., 96, 98, 100, ..., 255]`         | `[97, 98, 1, 99, 1]`            |
| 'b' (98)  | `[0, 99, 97, 1, 2, ..., 96, 98, 1, 100, ..., 255]`      | `[97, 98, 1, 99, 1, 2]`         |
| 'a' (97)  | `[0, 2, 99, 97, 1, 3, ..., 96, 98, 100, ..., 255]`      | `[97, 98, 1, 99, 1, 2, 1]`      |
| 'd' (100) | `[0, 2, 99, 97, 1, 3, ..., 96, 98, 100, 101, ..., 255]` | `[97, 98, 1, 99, 1, 2, 1, 100]` |

3. Result: The transformed sequence for `"abacabad"` is `[97, 98, 1, 99, 1, 2, 1, 100]`.
