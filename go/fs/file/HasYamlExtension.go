package file

const (
	yamlExtensionLong  = ".yaml"
	yamlExtensionShort = ".yml"
)

// HasYamlExtension - return boolean result if a YAML extension exists
func HasYamlExtension(filename string) bool {
	ext := GetExtension(filename)
	return ext == yamlExtensionLong || ext == yamlExtensionShort
}
