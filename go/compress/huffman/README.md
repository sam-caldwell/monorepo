Huffman Codec
=============
(c) 2024 Sam Caldwell. MIT License.

## Objective

* Develop a common huffman encoding/decoding package.

## Discussion

### What is Huffman Coding?

Huffman coding is a 1952 lossless data compression algorithm by David A. Huffman, which assigns variable-length
codes (Prefix Codes) to input characters. The length of these variable-length Prefix Codes is based on the frequencies
of corresponding characters.

The variable-length Prefix Codes are bit sequences assigned in a way that the code assigned to one character is unique
to that character. This means once a Prefix Code is assigned to one character, it cannot be assigned to any other
character. For example:

> Assume there exists four characters: a,b,c,d with corresponding Prefix Codes 00, 01, 0 and 1. This would fail to
> satisfy our rule because the code for 'c' (0) exists within the prefix codes for 'a' and 'b' (00 and 01). Thus, if
> a compressed bit stream is 00010, the de-compressed output may be 'cccdc' (0 0 0 1 0), 'abc' (00 01 0) or 'acdc'
> (00 0 1 0).

The common use case for huffman coding is data compression.

At the heart of the Huffman coding algorithm is the Huffman tree. The Huffman tree is a bottom-up frequency-sorted
binary tree.
---

### Algorithm: (See [`Encode()`](./Encode.go)

#### Building the Huffman Tree

1. We start building the huffman tree by splitting the entire input signal into
   tokens (e.g. bytes) we call "symbols.
2. These symbols are compiled into a [frequency table](./FrequencyTable.go) which
   we will use to build a frequency-sorted binary tree.

##### Frequency Table Calculation

1. The [frequency table](./FrequencyTable.go) is a map of `byte` (symbol) and an unsigned
   integer (`uint`) counter.
2. The [`Calculate()`](./FrequencyTable.Calculate.go) method iterates over the symbols
   in the input `[]byte` string and increments the counter.

##### Min-Heap (PriorityQueue) Creation

1. When [frequency table](./FrequencyTable.go) returns, the
   [`BuildHuffmanTree()`](./FrequencyTable.BuildHuffmanTree.go) is called.
2. This method calls [`LoadPriorityQueue()`](./FrequencyTable.LoadPriorityQueue.go) method to--
    1. Create a [`Node`](./Node.go) object.
    2. Set the [`Node`](./Node.go) `symbol` and `frequency` properties from the [frequency table](./FrequencyTable.go)
    3. Push the [`Node`](./Node.go) to a `PriorityQueue` heap structure.
3. We can then iterate over the priority queue to construct the tree.

##### Tree Construction

1. Using the `PriorityQueue`, [`BuildHuffmanTree()`](./FrequencyTable.BuildHuffmanTree.go) will--
    1. Repeatedly extract the two nodes with the smallest frequencies.
    2. Create a new parent node with these two nodes as children, and its frequency
       as the sum of their frequencies.
    3. Insert this new parent node back into the min heap.
    4. Continue this process until there's only one node left in the heap, which
       becomes the root of the Huffman tree.

> ***NOTE:***
>
>  Natively, Huffman can produce more than one `huffman tree` where two nodes have the same
> frequency. This is not an issue when compressing a signal since the `huffman codes` dictionary
> is included with the compressed signal and decoding from that codebook eliminates any issue. BUT
> the collisions do make testing difficult.
>
>  To overcome this problem, our implementation here is opinionated, and the
> [`PriorityQueue.Less()`](./PriorityQueue.Less.go) method first compares frequency. But where the frequency
> two nodes are equal, a tie-breaker comparing symbol values reduces the risk of a collision.

The following is an example tree, produced by [`Node.PrettyPrint()`](./Node.PrettyPrint.go):

```text

(root)
  │           ┌── (e:16)
  │       ┌── (30)
  │       │   │   ┌── (b:9)
  │       │   └── (14)
  │       │       └── (a:5)
  │   ┌── (55)
  │   │   │   ┌── (d:13)
  │   │   └── (25)
  │   │       └── (c:12)
  └── (100)
      └── (f:45)
```

#### Traversing the Tree to Assign Codes:

1. Given a `Huffman Tree` (depicted above) we can traverse the key by calling [`traverseTree()`](./traverseTree.go).
2. This function moves from top, down and left, right, assigning the left node a `0` and right node `1`.
3. The result is a map of symbols to unique binary codes:

```text
 map[97:[1 1] 98:[0 1] 99:[1 0 0 0] 100:[1 0 1] 101:[1 0 0 1] 102:[0 0]]
```

or--

| symbol | Huffman Code |
|--------|--------------|
| 97     | 11           |
| 98     | 01           |
| 99     | 1000         |
| 100    | 101          |
| 101    | 1001         |
| 202    | 00           |

---

### Algorithm: [`Decode()`](./Decode.go))

1. Given an `encodedInput` (`[]byte`) and `CodeMap`, we will--
    1. reverse the `CodeMap` order
    2. traverse the `encodedInput` to decode it.
2. If there are any leftover bits, we throw an error.
