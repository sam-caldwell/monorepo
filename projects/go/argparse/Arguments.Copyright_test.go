package argparse

import (
	"fmt"
	"testing"
)

func TestArguments_Copyright(t *testing.T) {
	var arg Arguments

	const (
		year   = 2023
		author = "Wendell Fertig"
		email  = "wfertig@philipines.tld"
	)

	expected := fmt.Sprintf("(c) %d %s <%s>",
		year, author, email)

	arg.Copyright(year, author, email)

	if text := arg.postscriptText.Pop(); text != expected {
		t.Logf("expected: '%s'", expected)
		t.Logf("    text: '%s'", text)
		t.Fatal("copyright mismatch")
	}
}
