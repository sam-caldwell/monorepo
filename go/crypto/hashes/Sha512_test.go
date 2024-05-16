package hashes

import (
	"crypto/sha512"
	"reflect"
	"testing"
)

func TestSha512(t *testing.T) {
	input := "hello"
	expectedHash := "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
	hash := Sha512(input)

	if hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, hash)
	}
}

func TestSha512Bytes(t *testing.T) {
	input := []byte("hello")
	expectedHash := sha512.Sum512(input)
	hash := Sha512Bytes(input)

	if !reflect.DeepEqual(hash, expectedHash[:]) {
		t.Errorf("Expected hash %x, got %x", expectedHash, hash)
	}
}

func TestSha512EmptyString(t *testing.T) {
	input := ""
	expectedHash := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
	hash := Sha512(input)

	if hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, hash)
	}
}

func TestSha512BytesEmptySlice(t *testing.T) {
	var input []byte
	expectedHash := sha512.Sum512(input)
	hash := Sha512Bytes(input)

	if !reflect.DeepEqual(hash, expectedHash[:]) {
		t.Errorf("Expected hash %x, got %x", expectedHash, hash)
	}
}
