package projectFilter

import (
	keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	"strings"
)

// FilterOnLanguage - given expectedLanguage, read MANIFEST.txt (if found); determine if project meets the expectation
func FilterOnLanguage(manifest keyvalue.KeyValue, expectedLanguage Filter) bool {
	if value, found := manifest.FindKey("language"); found {
		var detectedLanguage Filter
		detectedLanguage.FromString(strings.TrimSpace(strings.ToLower(value.(string))))
		return detectedLanguage == expectedLanguage
	}

	return expectedLanguage == LanguageGo //Default to golang
}
