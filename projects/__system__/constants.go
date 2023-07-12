package monorepo

import "strings"

const SupportedLanguages = "amd64Asm,arm64Asm,c,cpp,go,node,python,rust,typescript"

func GetSupportedLanguages() []string {
	return strings.Split(SupportedLanguages, ",")
}
