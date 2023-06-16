package ordered

func (set *Set) Pop() any {
	//popping item from set
	defer func() {
		if len(set.data) > 0 {
			set.data = set.data[1:]
		}
	}()
	return set.data[0]
}
