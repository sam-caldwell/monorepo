//go:build darwin
// +build darwin

package systemrecon

import (
	"testing"
)

func TestGetCacheSizes(t *testing.T) {
	//testCases := []struct {
	//	name       string
	//	level      int
	//	output     string
	//	expectedSz int
	//	expectErr  bool
	//}{
	//	{
	//		name:       "L1 Cache",
	//		level:      1,
	//		output:     "32768\n",
	//		expectedSz: 32,
	//		expectErr:  false,
	//	},
	//	{
	//		name:       "L2 Cache",
	//		level:      2,
	//		output:     "65536\n",
	//		expectedSz: 64,
	//		expectErr:  false,
	//	},
	//	{
	//		name:       "L3 Cache",
	//		level:      3,
	//		output:     "1048576\n",
	//		expectedSz: 1024,
	//		expectErr:  false,
	//	},
	//	{
	//		name:       "Invalid level",
	//		level:      4,
	//		output:     "",
	//		expectedSz: invalidCacheSz,
	//		expectErr:  true,
	//	},
	//	// Add more test cases as needed.
	//}
	//
	//for _, tc := range testCases {
	//	t.Run(tc.name, func(t *testing.T) {
	//		mockExecutor := runcommand.SimpleMockCommandExecutor{
	//			Output: tc.output,
	//			Error:  nil,
	//		}
	//
	//		got, err := getCacheSizes(mockExecutor, tc.level)
	//		if (err != nil) != tc.expectErr {
	//			t.Fatalf("Test %s error = %v, wantErr %v", tc.name, err, tc.expectErr)
	//			return
	//		}
	//		if got != tc.expectedSz {
	//			t.Fatalf("Test %s  data = %v, want %v", tc.name, got, tc.expectedSz)
	//		}
	//	})
	//}
}
