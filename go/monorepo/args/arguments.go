package args

type Arguments struct {
	option struct {
		debug bool
		color bool
		noop  bool
	}
	group      string
	command    string
	parameters map[string]string
}
