package findCollisions

// Finding - A finding is a report object describing some outcome of the AsynchronousJob()
type Finding struct {
	Id        int
	Collision bool
	Err       error
	Hash      string
	Raw       string
}
