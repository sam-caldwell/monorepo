package crsce

/*
 *
 */

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
	"testing"
)

func TestArguments_GetArgs(t *testing.T) {
	var arg Arguments
	func() {
		inFile, err := os.CreateTemp("", "testFileNameIn.txt")
		if err != nil {
			t.Fatal(err)
		}
		outFile, err := os.CreateTemp("", "testFileNameOut.txt")
		if err != nil {
			t.Fatal(err)
		}
		oName := outFile.Name()
		_ = outFile.Close()
		_ = file.Delete(oName)

		os.Args = os.Args[0:]
		os.Args = append(os.Args, "--in")
		os.Args = append(os.Args, inFile.Name())
		os.Args = append(os.Args, "--out")
		os.Args = append(os.Args, oName)
		arg.GetArgs()
		if arg.In != inFile.Name() {
			t.Fatal("--in mismatch")
		}
	}()
}
