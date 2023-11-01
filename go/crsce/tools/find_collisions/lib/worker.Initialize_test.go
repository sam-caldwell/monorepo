package hashScanner

import "testing"

/*
 * hashScanner worker.Initialize() test
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * hashScanner initializer test function
 */

func TestWorker_Initialize(t *testing.T) {
	t.Skipf("Incomplete")
	var w Worker
	if err := w.Initialize(0, 0, 1); err != nil {
		t.Fatal(err)
	}
	//ToDo: fix this test.
}
