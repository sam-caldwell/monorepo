package monorepo

// Config - Monorepo manifest structure
type Config struct {
	Header  HeaderSection `yaml:"header"`
	Clean   Stage         `yaml:"clean"`
	Build   Stage         `yaml:"build"`
	Install Stage         `yaml:"install"`
	Test    Stage         `yaml:"test"`
}
