package persistence

func (qry *QueryCollector) Count() uint64 {
	return qry.count.Current()
}
