package keyvault

// GetVersion - Return the keyvault major and minor version.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyVault) GetVersion() (uint8, uint8) {
	return kv.Version.Major, kv.Version.Minor
}
