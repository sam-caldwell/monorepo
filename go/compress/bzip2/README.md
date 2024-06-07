bzip2 Compression
=================

## Objectives
Implement a bzip2 `compress()` and `decompress()` function which will perform a
bzip2 compression on a given input `[]byte` slice.

---

## Usage
GIVEN `input []byte`, 
EXPECT `Compress(input) []byte` will compress `input` to a compressed `[]byte`.

GIVEN `input []byte`,
EXPECT `Decompress(input) []byte` will decompress `input` to `[]byte`.

---

## Algorithm (Compression)
### [Burrows-Wheeler Transform (BWT)](./bzip2bwt/Bwt.go):
The input data is transformed using the Burrows-Wheeler Transform. This 
rearranges the data into runs of similar characters, which is easier to 
compress.

### [Move-to-Front Transform (MTF)](./bzip2mtf/Mtf.go):
The BWT output is then processed with the Move-to-Front transform. This step 
converts repeated characters into sequences of smaller numbers, enhancing the 
effectiveness of subsequent steps.

### [Run-Length Encoding (RLE)](../rle/Encode.go):
The MTF output is further compressed using Run-Length Encoding. This reduces 
the size of long runs of the same character.

### Huffman Coding:
Finally, the data is compressed using Huffman coding, which replaces 
frequently occurring symbols with shorter codes and less frequent symbols 
with longer codes.

---
## Algorithm (Decompression)
### Huffman Decoding:

The compressed data is first decoded using Huffman decoding to reconstruct 
the Run-Length Encoded data.

### Run-Length Decoding (RLD):

The output from the Huffman decoding is then processed to revert the 
Run-Length Encoding, restoring the sequences of repeated characters.

### Inverse Move-to-Front Transform (IMTF):

The RLD output undergoes the Inverse Move-to-Front transform, converting 
sequences of numbers back into the original character sequences.

### Inverse Burrows-Wheeler Transform (IBWT):

Finally, the data is restored to its original form using the Inverse 
Burrows-Wheeler Transform.
