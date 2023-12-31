package semver

const (
	//defaultRepo indicates the current local directory
	defaultRepo = "."

	defaultTag = "v0.0.0"

	tagPattern = `^[vV][0-9]+.[0-9]+.[0-9]+$`

	tagFormat = "v%d.%d.%d"

	dotDelimiter = "."

	// This is the number of bits in each part of a semantic version (0-255)
	versionNumberSize = 8

	versionNumberBase = 10

	errVersionOverflow = "version overflow"

	errNilPointer       = "nil pointer"
	errTagAlreadyExists = "tag '%s' already exists"
)
