package monorepo

// HeaderSection - Monorepo manifest header
type HeaderSection struct {
	Author      string   `yaml:"author"`
	Email       string   `yaml:"email"`
	Description string   `yaml:"description"`
	Arch        []string `yaml:"arch"`
	Opsys       []string `yaml:"opsys"`
}
