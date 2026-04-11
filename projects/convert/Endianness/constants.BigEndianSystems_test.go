package endianness

import (
	"runtime"
	"testing"
)

//goland:noinspection GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions,GoBoolExpressions
func TestEndianness(t *testing.T) {
	/*
	 * Test plan...
	 * 1. Evaluate goarch to determine if we have a big or little endian architecture, creating two bools.
	 */

	littleEndian := (
	// Any of these systems are little endian
	runtime.GOARCH == "amd64" ||
		runtime.GOARCH == "arm64" ||
		runtime.GOARCH == "386" ||
		runtime.GOARCH == "arm" ||
		runtime.GOARCH == "mips" ||
		runtime.GOARCH == "mipsle" ||
		runtime.GOARCH == "mips64" ||
		runtime.GOARCH == "mips64le" ||
		runtime.GOARCH == "ppc64" ||
		runtime.GOARCH == "ppc64le" ||
		runtime.GOARCH == "riscv64" ||
		runtime.GOARCH == "s390x" ||
		runtime.GOARCH == "sparc64" ||
		runtime.GOARCH == "wasm" ||
		runtime.GOARCH == "wasm64") &&
		!(
		// Any of these systems are big endian
		runtime.GOARCH == "ppc" ||
			runtime.GOARCH == "ppc64" ||
			runtime.GOARCH == "s390" ||
			runtime.GOARCH == "s390x")

	if littleEndian {
		/*
		 * We are on a little endian system.  Make sure our constants are valid
		 */
		if !LittleEndian {
			t.Fatalf("Expect that LittleEndian would be true on %v", runtime.GOARCH)
		}
		if BigEndian {
			t.Fatalf("Expect that BigEndian would be false on %v", runtime.GOARCH)
		}
	} else {
		/*
		 * We are on a big endian system.  Make sure our constants are valid
		 */
		if !BigEndian {
			t.Fatalf("Expect that BigEndian would be true on %v", runtime.GOARCH)
		}
		if LittleEndian {
			t.Fatalf("Expect that LittleEndian would be false on %v", runtime.GOARCH)
		}
	}
	if BigEndian && LittleEndian {
		t.Fatal("BigEndian and LittleEndian cannot both be true")
	}
	if !BigEndian && !LittleEndian {
		t.Fatal("BigEndian and LittleEndian cannot both be false")
	}
}
