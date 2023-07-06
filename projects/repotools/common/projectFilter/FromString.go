package projectFilter

import "strings"

// FromString - Convert a string input to a corresponding filter number (or None if no match)
func (o *Filter) FromString(s string) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "hideenabled":
		*o = HideEnabled
	case "hidedisabled":
		*o = HideDisabled
	case "windows":
		*o = Windows
	case "linux":
		*o = Linux
	case "darwin":
		*o = Darwin
	case "command":
		*o = Command
	case "package":
		*o = Package
	case "languagego":
		*o = LanguageGo
	case "languagec":
		*o = LanguageC
	case "languagecpp":
		*o = LanguageCpp
	case "languageasmamd64":
		*o = LanguageAsmAmd64
	case "languageasmarm64":
		*o = LanguageAsmArm64
	case "languagerust":
		*o = LanguageRust
	case "languagenodejs":
		*o = LanguageNodeJs
	case "languagepython":
		*o = LanguagePython
	default:
		*o = None
	}
}
