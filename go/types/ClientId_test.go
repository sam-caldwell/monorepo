package types

import (
	"github.com/google/uuid"
	"testing"
)

func TestClientId(t *testing.T) {
	var c ClientId
	u, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	c = ClientId(u)
	if u.String() != c.String() {
		t.Fatal("uuid->clientID convert fail")
	}
}
