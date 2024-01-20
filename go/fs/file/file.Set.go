package file

// Set - Set the given file/path name to the internal state (no validation)
func (fp *File) Set(name string) error {
	if err := fp.valid(&name); err != nil {
		*fp = ""
		return err
	}
	*fp = File(name)
	return nil
}
