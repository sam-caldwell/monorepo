package list

// Add - Add a record to the list (list must be empty or match type of list[0])
func (list *SmartList) Add(record any) (err error) {
	if err = list.TypeCheck(record); err == nil {
		*list = append(*list, record)
	}
	return err
}
