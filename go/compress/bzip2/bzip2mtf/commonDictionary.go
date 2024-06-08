package bzip2mtf

// commonDictionary - the common dictionary is an agreed pattern shared across Mtf() and Imtf()
//
// (c) 2023 Sam Caldwell.  MIT License
var commonDictionary [256]byte

// init - we initialize the dictionary once to be used for all MTF/iMTF calls.
//
// (c) 2023 Sam Caldwell.  MIT License
func init() {
	// Initialize common dictionary with values 0 to 255
	for i := 0; i < 256; i++ {
		commonDictionary[i] = byte(i)
	}
}
