package directory

import "fmt"

// Path - A strongly typed directory.Path with methods
type Path string

// NewPath - Return a new path with a given string or []byte value.
func NewPath[T string | []byte](p T, mustExist bool) (*Path, error) {
	var o Path
	if mustExist {
		if !Exists(string(p)) {
			return nil, fmt.Errorf("path must exist but not found")
		}
	}
	o.Set(string(p))
	return &o, nil
}

// String - Getter: return string path
func (p *Path) String() string {
	return (string)(*p)
}

// Bytes - Getter: return a []byte path
func (p *Path) Bytes() []byte {
	return ([]byte)(*p)
}

// Exists - Return whether the path exists.
func (p *Path) Exists() bool {
	return Exists(string(*p))
}

// Setp - Setter: store a new string (*v) to the Path.
func (p *Path) Setp(v *string) {
	//ToDo: validate path is valid/sanitized string
	*p = Path(*v)
}

// Set - Setter: store a new string (v) to the Path.
func (p *Path) Set(v string) {
	p.Setp(&v)
}
