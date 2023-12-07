package monorepo

type Config struct {
	Clean   Stage `yaml:"clean"`
	Build   Stage `yaml:"build"`
	Install Stage `yaml:"install"`
	Test    Stage `yaml:"test"`
}
