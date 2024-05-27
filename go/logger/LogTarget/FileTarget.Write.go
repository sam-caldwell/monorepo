package LogTarget

// Write - Write a formatted log string to stdout.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (out FileTarget) Write(p *[]byte) error {
	_, err := out.file.WriteBytes(*p)
	return err
}
