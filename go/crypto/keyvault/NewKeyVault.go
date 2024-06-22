package keyvault

// NewKeyVault - Create a new version 0.0 KeyVault object
//
//	(c) 2023 Sam Caldwell.  MIT License
func NewKeyVault(fileName string) *KeyVault {
	return &KeyVault{
		fileName: fileName,
		Version: VersionStruct{
			Major: 0,
			Minor: 0,
		},
	}
}
