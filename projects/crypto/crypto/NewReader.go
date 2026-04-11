package crypto

import (
    "crypto/rsa"
)

//  Theory of operation
//    1. A key exists as a public and private key.
//    2. The public key is written to a public-read s3 object
//    3. The private key is encrypted by a symmetric encryption on a private s3 object



// Rsa - Manage the keys and encryption lifecycle.
//
//		Imagine an application has an encrypted file stored in an s3 bucket. This file is first encrypted
//		with a symmetric key using a passphrase known only to the app (and possibly user) then with an
//		asymmetric key known only to the application.
//
//      Key rotation:
//          1. Keys are rotated by adding a new key pair and deleting the public key from the previously
//      current key.
//          2. A key is rotated off the list
//      Keys are rotated by adding them to the `keys` map.  When a key is replaced, they key removes the
//      public key to prevent it from encrypting any information, so only the newest key can be used.
//
//
type Rsa struct {
    mostRecentKey uint
    private       rsa.PrivateKey
    public        []rsa.PublicKey
}

// NewRsaManager - Return a Rsa Encryption manager reference
func NewRsaManager(s3bucket, keyFile *string) (out *Rsa, err error) {
    var o Rsa
    if err = o.LoadKeys(s3bucket, keyFile); err != nil {
        return nil, err
    }
    return out, err
}

// LoadKeys - Load the encryption keys file into the manager's memory.
func (o *Rsa) LoadKeys() (, err error) {

return out, err
}

// UpdateKeys - Update
