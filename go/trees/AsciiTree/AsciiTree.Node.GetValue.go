package asciitree

import "fmt"

// GetValue - Return the Node value, formatted properly.
func (n *Node) GetValue() string {
	if o, ok := n.Value.(string); ok {
		return fmt.Sprintf("%s", o)
	}
	if o, ok := n.Value.([]string); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.(float32); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.([]float32); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.(float64); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.([]float64); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.(bool); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.([]bool); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.(byte); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.([]byte); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := n.Value.(rune); ok {
		return fmt.Sprintf("%c", o)
	}
	if o, ok := n.Value.([]rune); ok {
		return fmt.Sprintf("%v", o)
	}

	if o, ok := n.Value.(int); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(int8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(int16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(int32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(int64); ok {
		return fmt.Sprintf("%d", o)
	}

	if o, ok := n.Value.(uint); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(uint8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(uint16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(uint32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.(uint64); ok {
		return fmt.Sprintf("%d", o)
	}

	if o, ok := n.Value.([]int); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]int8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]int16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]int32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]int64); ok {
		return fmt.Sprintf("%d", o)
	}

	if o, ok := n.Value.([]uint); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]uint8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]uint16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]uint32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := n.Value.([]uint64); ok {
		return fmt.Sprintf("%d", o)
	}
	return fmt.Sprintf("%v", n.Value)
}
