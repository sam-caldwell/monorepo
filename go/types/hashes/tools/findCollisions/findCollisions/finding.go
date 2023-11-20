package find_collisions

// Finding - A finding is a report object describing some outcome of the AsynchronousJob()
type Finding struct {
	id        int
	collision bool
	err       error
	hash      string
	raw       string
}
