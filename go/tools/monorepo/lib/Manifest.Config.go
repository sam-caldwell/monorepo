package monorepo

type Config struct {
	Clean Stage `yaml:"clean"`
	Build Stage `yaml:"build"`
	Test  Stage `yaml:"test"`
}
