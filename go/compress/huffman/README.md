Huffman Codec
=============

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

### Algorithm: Encoding

#### Building the Huffman Tree

##### Frequency Table Calculation

Calculate the frequency of each symbol in the input data.

```go
package main

// CalculateFrequencyTable takes a byte slice and returns a map with the frequency of each byte.
func CalculateFrequencyTable(s []byte) map[byte]uint {
    table := make(map[byte]uint)
    for _, c := range s {
        table[c]++
    }
    return table
}
```

##### Min Heap Creation

Create a min heap (priority queue) with a leaf node for each unique symbol based on their frequencies:

1. Define a Node struct to represent each leaf node.

```go
package main

// Node represents a leaf node with a symbol and its frequency
type Node struct {
    symbol    byte
    frequency uint
}
```

2. Implement the heap.Interface for a priority queue.

```go
package main

// PriorityQueue implements the heap.Interface and holds Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want the lower frequency to be the higher priority
    return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    // Push the element onto the heap
    *pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}
```

3. Create the priority queue and insert nodes based on their frequencies.

```go
package main

func main() {
    // Assume frequencyTable is calculated as defined earlier in this
    // documentation...

    // Create a priority queue and initialize it
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)

    // Insert each symbol and its frequency as a Node into the priority queue
    for symbol, frequency := range frequencyTable {
        heap.Push(&pq, &Node{symbol: symbol, frequency: frequency})
    }
}
```

##### Tree Construction

1. Repeatedly extract the two nodes with the smallest frequencies.
2. Create a new parent node with these two nodes as children, and its frequency as the sum of their frequencies.
3. Insert this new parent node back into the min heap.
4. Continue this process until there's only one node left in the heap, which becomes the root of the Huffman tree.

```go
package main

// Assume the above sections define 
// the frequency table, priorityQueue/heap, etc....

func buildHuffmanTree(frequencyTable map[byte]uint) *Node {
    // Create a priority queue and initialize it
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)

    // Insert each symbol and its frequency as a Node into the priority queue
    for symbol, frequency := range frequencyTable {
        heap.Push(&pq, &Node{symbol: symbol, frequency: frequency})
    }

    // Tree construction process
    for pq.Len() > 1 {
        // Step 1: Repeatedly extract the two nodes with the smallest frequencies
        left := heap.Pop(&pq).(*Node)
        right := heap.Pop(&pq).(*Node)

        // Step 2: Create a new parent node with these two nodes as children,
        // and its frequency as the sum of their frequencies
        parent := &Node{
            frequency: left.frequency + right.frequency,
            left:      left,
            right:     right,
        }

        // Step 3: Insert this new parent node back into the min heap
        heap.Push(&pq, parent)
    }

    // Step 4: Continue this process until there's only one node left in the heap,
    // which becomes the root of the Huffman tree
    return heap.Pop(&pq).(*Node)
}
```

#### Traversing the Tree to Assign Codes:

1. Traverse the Huffman tree to assign a unique binary code to each symbol.
2. Typically, a left edge represents a '0' and a right edge represents a '1'.
3. The final output is a `map` of unique binary codes for each symbol, which can be used for efficient data encoding.

```go
package main

//
// Assume the above sections define the referenced functions, etc....
//

// traverseTree assigns binary codes to each symbol by traversing the Huffman tree
func traverseTree(root *Node, prefix string, codes map[byte]string) {
    if root == nil {
        return
    }

    if root.left == nil && root.right == nil {
        codes[root.symbol] = prefix
        return
    }

    traverseTree(root.left, prefix+"0", codes)
    traverseTree(root.right, prefix+"1", codes)
}
```

#### Output Codes

The algorithm will output a `map` of symbols (e.g. bytes) and unique Prefix Codes:

```go
package main

func main() {
    s := "my input string"
    frequencyTable := CalculateFrequencyTable([]byte(s))
    root := buildHuffmanTree(frequencyTable)

    codes := make(map[byte]string)
    traverseTree(root, "", codes)

    // Print the Huffman codes for each symbol
    for symbol, code := range codes {
        fmt.Printf("Symbol: %c, Code: %s\n", symbol, code)
    }
}
```

---

### Algorithm: Decoding
#### Summary
1. Start from the root of the Huffman tree.
2. For each bit in the encoded string, traverse the tree:
    1. Move to the left child if the bit is '0'.
    2. Move to the right child if the bit is '1'.
    3. When a leaf node is reached, output the symbol and return to the root to decode the next sequence.
    4. Repeat this process until the entire encoded string is decoded.
#### Example Code
```go
package main

// decodeHuffman takes the Huffman tree root and an encoded string,
// and returns the decoded original string
func decodeHuffman(root *Node, encodedStr string) string {
    var decodedStr string
    currentNode := root

    for _, bit := range encodedStr {
        if bit == '0' {
            currentNode = currentNode.left
        } else {
            currentNode = currentNode.right
        }

        // If we reach a leaf node
        if currentNode.left == nil && currentNode.right == nil {
            decodedStr += string(currentNode.symbol)
            currentNode = root // Go back to the root for the next symbol
        }
    }

    return decodedStr
}
```
#### Test Code
```go
package main

func main() {
    inputString := "my input string"
    frequencyTable := CalculateFrequencyTable([]byte(inputString))
    root := buildHuffmanTree(frequencyTable)

    codes := make(map[byte]string)
    traverseTree(root, "", codes)

    // Encode the input string
    var encodedStr string
    for _, c := range s {
        encodedStr += codes[byte(c)]
    }

    fmt.Println("Encoded String:", encodedStr)

    // Decode the encoded string
    decodedString := decodeHuffman(root, encodedStr)
    fmt.Println("Decoded String:", decodedString)

    if inputString != decodedString {
        fmt.Println("The algorithm has failed.  input should have matched the decoded result.")
    }
}
```
