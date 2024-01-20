package file

// Get - Return the internal file/path state
func (fp *File) Get() string {
	return string(*fp)
}
