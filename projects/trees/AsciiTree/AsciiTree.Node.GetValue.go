package asciitree

import "fmt"

// GetValue - Return the Node value, formatted properly.
func (node *Node) GetValue() string {
	if o, ok := node.Value.(string); ok {
		return fmt.Sprintf("%s", o)
	}
	if o, ok := node.Value.([]string); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.(float32); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.([]float32); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.(float64); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.([]float64); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.(bool); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.([]bool); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.(byte); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.([]byte); ok {
		return fmt.Sprintf("%v", o)
	}
	if o, ok := node.Value.(rune); ok {
		return fmt.Sprintf("%c", o)
	}
	if o, ok := node.Value.([]rune); ok {
		return fmt.Sprintf("%v", o)
	}

	if o, ok := node.Value.(int); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(int8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(int16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(int32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(int64); ok {
		return fmt.Sprintf("%d", o)
	}

	if o, ok := node.Value.(uint); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(uint8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(uint16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(uint32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.(uint64); ok {
		return fmt.Sprintf("%d", o)
	}

	if o, ok := node.Value.([]int); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]int8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]int16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]int32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]int64); ok {
		return fmt.Sprintf("%d", o)
	}

	if o, ok := node.Value.([]uint); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]uint8); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]uint16); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]uint32); ok {
		return fmt.Sprintf("%d", o)
	}
	if o, ok := node.Value.([]uint64); ok {
		return fmt.Sprintf("%d", o)
	}
	return fmt.Sprintf("%v", node.Value)
}
