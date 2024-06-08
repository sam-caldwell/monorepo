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

1. Assume an input sequence `"abacabad"`
2. The MTF list is initialized to `[0, 1, 2, ..., 255]`.
3. The following table shows the MTF process:

| Input             | Output | Notes                                                                 |
|-------------------|--------|-----------------------------------------------------------------------|
| `a` (ASCII `97`)  |        | `a` is at index `97` in the initial list.                             |
|                   | `97`   | Move a to the front: `[a, 0, 1, 2, ..., 96, 98, ..., 255]`            |
| `b` (ASCII `98`)  |        | `b` is at index `98` in the updated list.                             |
|                   | `98`   | Move b to the front: `[b, a, 0, 1, 2, ..., 96, 98, ..., 255]`         |
| `a`               |        | `a` is now at index `1`.                                              |
|                   | `1`    | Move a to the front: `[a, b, 0, 1, 2, ..., 96, 98, ..., 255]`         |
| `c` (ASCII `99`)  |        | `c` is at index `99` in the updated list.                             |
|                   | `99`   | Move c to the front: `[c, a, b, 0, 1, 2, ..., 96, 98, ..., 255]`      |
| `a`               |        | `a` is at index `1`.                                                  |
|                   | `1`    | Move a to the front: `[a, c, b, 0, 1, 2, ..., 96, 98, ..., 255]`      |
| `b`               |        | `b` is at index `2`.                                                  |
|                   | `2`    | Move b to the front: `[b, a, c, 0, 1, 2, ..., 96, 98, ..., 255]`      |
| `a`               |        | `a` is at index `1`.                                                  |
|                   | `1`    | Move a to the front: `[a, b, c, 0, 1, 2, ..., 96, 98, ..., 255]`      |
| `d` (ASCII `100`) |        | `d` is at index `100` in the updated list.                            |
|                   | `100`  | Move `d` to the front: `[d, a, b, c, 0, 1, 2, ..., 96, 98, ..., 255]` |

4. Result: The transformed sequence for `"abacabad"` is `[97, 98, 1, 99, 1, 2, 1, 100]`.
