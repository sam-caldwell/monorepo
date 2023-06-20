//go:build linux
// +build linux

package systemrecon

func GetVersion() (version string, err error) {
	var bytes []byte
	distro, err := GetLinuxDistro()
	if err != nil {
		return version, err
	}
	switch distro {
	case words.Debian:
		bytes, err = os.ReadFile(DebianVersion)
		version = string(bytes)
	case words.Ubuntu:
		bytes, err = os.ReadFile(DebianVersion)
		version = string(bytes)
	}
	if version == "" {
		err = fmt.Errorf(exit.ErrUnsupportedOpsys)
	}
	return version, err
}
