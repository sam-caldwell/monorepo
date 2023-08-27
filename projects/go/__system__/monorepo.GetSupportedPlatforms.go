package monorepo

import "strings"

func GetSupportedPlatforms() []string {
	return strings.Split(SupportedPlatforms, ",")
}
