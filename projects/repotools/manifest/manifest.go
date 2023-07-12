package projectmanifest

type Manifest struct {
	fileName           string
	Name               string              `yaml:"name"`
	Author             string              `yaml:"author"`
	Options            ManifestOptions     `yaml:"options"`
	SupportedPlatforms []ManifestPlatforms `yaml:"supportedPlatforms"`
	Dependencies       []string            `yaml:"dependencies"`
	err                error
}
