package misc

import "testing"

func TestRightPad(t *testing.T) {
	const (
		space      = ' '
		underscore = '_'
		root       = "test"
	)

	var expectedPadding [2]string
	pad := [2]rune{space, underscore}

	for i := len(root) + 1; i <= 10*len(root); i++ {
		expectedPadding[0] += string(space)
		expectedPadding[1] += string(underscore)

		for n := 0; n < 2; n++ {
			expected := root + expectedPadding[n]

			if output := RightPad(pad[n], root, i); (output != expected) || (len(output) != i) {
				t.Fatalf("output[%d]: '%s'(%d) != '%s'(%d)", i, output, len(output), expected, len(expected))
			}
		}
	}
	if output := RightPad(space, root, len(root)); output != root {
		t.Fatal("same length test failed")
	}
	if output := RightPad(space, root, 0); output != root {
		t.Fatal("0 length test failed")
	}

	for i := 0; i < 2*len(root); i++ {
		if output := RightPad(space, root, len(root)-i); output != root {
			t.Fatal("negative length test failed")
		}
	}
}
