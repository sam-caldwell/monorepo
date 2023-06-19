package ordered

func (set *Set) Pop() any {
	if len(set.data) > 0 {
		set.lock.Lock()
		//popping item from set
		defer func() {
			defer set.lock.Unlock()
			if len(set.data) > 0 {
				set.data = set.data[1:]
			}
		}()
		return set.data[0]
	}
	return nil
}
