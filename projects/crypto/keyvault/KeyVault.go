package keyvault

type KeyVault struct {
	fileName string
	Version  VersionStruct
	Payload  CipherPayload
}
