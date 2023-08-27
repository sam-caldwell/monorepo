package monorepo

import "strings"

func GetSupportedLanguages() []string {
	return strings.Split(SupportedLanguages, ",")
}
