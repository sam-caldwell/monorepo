package projectmanifest

type ManifestOptions struct {
	BuildEnabled bool   `yaml:"buildEnabled"`
	LintEnabled  bool   `yaml:"lintEnabled"`
	PackEnabled  bool   `yaml:"packEnabled"`
	ScanEnabled  bool   `yaml:"scanEnabled"`
	SignEnabled  bool   `yaml:"signEnabled"`
	Language     string `yaml:"language"`
}
