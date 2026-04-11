package findCollision

// Finding - A finding reporting object
type Finding struct {
	Id        int
	Collision bool
	Err       error
	Hash      string
	Raw       string
}
