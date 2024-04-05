package net

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"log"
	"regexp"
	"strings"
)

// IsValidFqdn - reports whether host is a valid hostname that can be matched or
// matched against according to RFC 6125 2.2, with some leniency to accommodate
// legacy values.
func IsValidFqdn(s string) bool {
	// Ensure the length of the string is within the allowed range
	if len(s) < 1 || len(s) > 255 {
		log.Printf("Failed1")
		return false
	}

	// Split the string into labels (trim trailing dot)
	labels := strings.Split(strings.TrimSuffix(s, "."), ".")

	// Ensure the first and last labels meet the length requirements
	if len(labels[0]) < 1 || len(labels[len(labels)-1]) < 2 || len(labels[len(labels)-1]) > 63 {
		log.Printf("Failed3 on %s", labels)
		return false
	}

	if (labels[0] == words.Hyphen) || (labels[len(labels)-1] == words.Hyphen) {
		return false
	}

	// Regular expression for valid label characters
	validLabel := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)

	// Check if each label is valid
	for _, label := range labels {
		label = strings.TrimSuffix(label, ".")
		if !validLabel.MatchString(label) {
			log.Printf("Failed4 on %s", label)
			return false
		}
		if strings.Contains(label, "--") || strings.Contains(label, "..") {
			return false
		}
	}
	return true
}
