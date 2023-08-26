package pair

func (o *OrderedPair) Walk(fn func(key string, value interface{}) error) error {
	for _, record := range o.data {
		if err := fn(record.Key, record.Value); err != nil {
			return err
		}
	}
	return nil
}
