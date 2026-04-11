package keyvault

const (
	//textWidth - the number of characters in each line of base64-encoded output to the keyvault file.
	textWidth = 80

	//headerPattern - the string pattern used to prepend the base64-encoded output to the keyvault file.
	headerPattern = "===KEY VAULT(Version 0x%02x.0x%02x)START===\n"

	//footerPattern - the string pattern used to append the base64-encoded output to the keyvault file.
	footerPattern = "===KEY VAULT(Version 0x%02x.0x%02x)END===\n"
)
