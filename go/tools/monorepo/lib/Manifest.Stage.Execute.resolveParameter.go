package monorepo

import "strings"

// resolveParameter - resolve the parameters in a Manifest.yml line with the matching values (in memory)
func resolveParameter(parameters map[string]string, line string) string {
	for key, value := range parameters {
		line = strings.Replace(line, key, value, -1)
	}
	return line
}
