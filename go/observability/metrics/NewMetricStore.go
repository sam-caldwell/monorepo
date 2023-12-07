package metrics

func NewMetricsStore() *MetricStore {
	return &MetricStore{
		data: make(map[string]any),
		err:  make([]error, 0),
	}
}
