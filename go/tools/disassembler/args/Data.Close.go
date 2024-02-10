package args

func (data *Data) Close() {
	defer func() {
		if data.source != nil {
			_ = data.source.Close()
		}
	}()
	defer func() {
		if data.out != nil {
			_ = data.out.Close()
		}
	}()
}
