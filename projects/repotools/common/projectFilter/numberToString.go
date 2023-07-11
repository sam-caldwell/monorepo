package projectFilter

import "fmt"

// numberToString - Convert an evenly divisible ListProjectFilter value to a string
func numberToString(f Filter) (string, error) {
	switch f {
	case None:
		return "None", nil
	case HideEnabled:
		return "HideEnabled", nil
	case HideDisabled:
		return "HideDisabled", nil
	case Windows:
		return "Windows", nil
	case Linux:
		return "Linux", nil
	case Darwin:
		return "Darwin", nil
	case Command:
		return "Command", nil
	case Package:
		return "Package", nil
	case LanguageGo:
		return "LanguageGo", nil
	case LanguageC:
		return "LanguageC", nil
	case LanguageCpp:
		return "LanguageCpp", nil
	case LanguageAsmAmd64:
		return "LanguageAsmAmd64", nil
	case LanguageAsmArm64:
		return "LanguageAsmArm64", nil
	case LanguageRust:
		return "LanguageRust", nil
	case LanguageNodeJs:
		return "LanguageNodeJs", nil
	case LanguagePython:
		return "LanguagePython", nil
	default:
		return "None", fmt.Errorf("invalid filter")
	}

}
