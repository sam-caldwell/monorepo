//go:build darwin
// +build darwin

package systemrecon

import (
	"testing"
)

func TestCpuCache(t *testing.T) {

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			result, err := CpuCache(tc.Executor)
			if err != nil {
				if tc.ExpectedError != nil {
					if err.Error() != tc.ExpectedError.Error() {
						t.Fatalf("Expected error not actual error (%v) (%v)", tc.ExpectedError, err)
					}
					return
				}
				t.Fatalf("Unexpected error: %v", err)
			}
			if tc.ExpectedError != nil {
				t.Fatalf("Expected error (%v) but got none.", tc.ExpectedError)
			}
			if result != tc.ExpectedResult {
				t.Fatalf("Expected result (%v) but got (%v)", result, tc.ExpectedError)
			}
		})
	}
}
