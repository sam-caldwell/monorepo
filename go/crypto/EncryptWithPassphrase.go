package main

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"io"
	"strings"
)

// EncryptWithPassphrase - Encrypt a clearText string with a given passphrase.
func EncryptWithPassphrase(clearText, passphrase string) (cipher string, err error) {
	var output bytes.Buffer
	var encode io.WriteCloser
	var encrypted io.WriteCloser

	if strings.TrimSpace(passphrase) == words.EmptyString {
		return words.EmptyString, fmt.Errorf("invalid passphrase (cannot be empty or whitespace)")
	}
	if encode, err = armor.Encode(&output, "PGP MESSAGE", nil); err != nil {
		return "", err
	}
	defer func() {
		_ = encode.Close()
	}()

	// Encrypt the clear text symmetrically with the passphrase
	if encrypted, err = openpgp.SymmetricallyEncrypt(encode, []byte(passphrase), nil, nil); err != nil {
		return words.EmptyString, err
	}
	if _, err = encrypted.Write([]byte(clearText)); err != nil {
		return words.EmptyString, err
	}
	if err = encrypted.Close(); err != nil {
		return words.EmptyString, err
	}
	return output.String(), nil
}
