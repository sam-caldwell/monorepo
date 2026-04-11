package huffman

import "sort"

func Encode(freqs map[byte]int, codes map[byte][]bool) {
	tree := buildHuffmanTree(freqs)
	generateHuffmanCodes(tree.Root, []bool{}, codes)

}

type Node struct {
	Symbol byte
	Freq   int
	Left   *Node
	Right  *Node
}

type Tree struct {
	Root *Node
}

func buildHuffmanTree(freqs map[byte]int) *Tree {
	var nodes []*Node
	for sym, freq := range freqs {
		nodes = append(nodes, &Node{Symbol: sym, Freq: freq})
	}

	for len(nodes) > 1 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Freq < nodes[j].Freq
		})

		left := nodes[0]
		right := nodes[1]
		nodes = nodes[2:]

		parent := &Node{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		nodes = append(nodes, parent)
	}

	return &Tree{Root: nodes[0]}
}

func generateHuffmanCodes(node *Node, prefix []bool, codes map[byte][]bool) {
	if node.Left == nil && node.Right == nil {
		codes[node.Symbol] = prefix
		return
	}
	generateHuffmanCodes(node.Left, append(prefix, false), codes)
	generateHuffmanCodes(node.Right, append(prefix, true), codes)
}
