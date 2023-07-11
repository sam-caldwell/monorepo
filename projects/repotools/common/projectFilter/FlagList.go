package projectFilter

// FlagList - return a list of flags
func FlagList() []Filter {
	return []Filter{
		None, HideEnabled, HideDisabled, Windows, Linux, Darwin, Command, Package, LanguageGo, LanguageC, LanguageCpp,
		LanguageAsmAmd64, LanguageAsmArm64, LanguageRust, LanguageNodeJs, LanguagePython}
}
