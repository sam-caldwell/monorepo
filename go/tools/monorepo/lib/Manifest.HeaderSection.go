package monorepo

// HeaderSection - Monorepo manifest header
type HeaderSection struct {
	Author      string   `yaml:"author"`
	Email       string   `yaml:"email"`
	Description string   `yaml:"description"`
	Opsys       []string `yaml:"opsys"`
	Arch        []string `yaml:"arch"`
}
