package pair

import (
	"bytes"
	"encoding/binary"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"math"
	"testing"
)

func TestKeyToBytes(t *testing.T) {
	t.Run("test with boolean type", func(t *testing.T) {
		var n bool
		t.Run("boolean value test", func(t *testing.T) {
			t.Run("test with false", func(t *testing.T) {
				n = false
				v, err := KeyToBytes(n)
				if err != nil {
					t.Fatalf("expected error for boolean type but got %v", err)
				}
				if v[0] != 0x00 {
					t.Fatalf("expected 0x00 for boolean type but got %v", v)
				}
			})
			t.Run("test with true", func(t *testing.T) {
				n = true
				v, err := KeyToBytes(n)
				if err != nil {
					t.Fatalf("expected error for boolean type but got %v", err)
				}
				if v[0] != 0xFF {
					t.Fatalf("expected 0xFF for boolean type but got %v", v)
				}
			})
		})
		t.Run("boolean pointer test", func(t *testing.T) {
			t.Run("test with false", func(t *testing.T) {
				n = false
				v, err := KeyToBytes(&n)
				if err != nil {
					t.Fatalf("expected error for boolean type but got %v", err)
				}
				if v[0] != 0x00 {
					t.Fatalf("expected 0x00 for boolean type but got %v", v)
				}
			})
			t.Run("test with true", func(t *testing.T) {
				n = true
				v, err := KeyToBytes(&n)
				if err != nil {
					t.Fatalf("expected error for boolean type but got %v", err)
				}
				if v[0] != 0xFF {
					t.Fatalf("expected 0xFF for boolean type but got %v", v)
				}
			})
		})
	})
	t.Run("test with signed integer type", func(t *testing.T) {
		t.Run("test with int", func(t *testing.T) {
			for expected := range []int{math.MinInt, -1024, -1, 0, 1, 1024, math.MaxInt} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with int8", func(t *testing.T) {
			for expected := range []int8{-127, -1, 0, 1, 127} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with int16", func(t *testing.T) {
			for expected := range []int16{math.MinInt16, -1024, -1, 0, 1, 1024, math.MaxInt16} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with int32", func(t *testing.T) {
			for expected := range []int32{math.MinInt32, -1024, -1, 0, 1, 1024, math.MaxInt32} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with int64", func(t *testing.T) {
			for expected := range []int64{math.MinInt64, -1024, -1, 0, 1, 1024, math.MaxInt64} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
	})
	t.Run("test with unsigned integer type", func(t *testing.T) {
		t.Run("test with uint", func(t *testing.T) {
			for expected := range []uint{0, 1, 1024, math.MaxUint} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with uint8", func(t *testing.T) {
			for expected := range []uint8{0, 1, math.MaxUint8} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with uint16", func(t *testing.T) {
			for expected := range []uint16{0, 1, math.MaxUint16} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with uint32", func(t *testing.T) {
			for expected := range []uint32{0, 1, math.MaxUint32} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})
		t.Run("test with uint64", func(t *testing.T) {
			for expected := range []uint64{0, 1, math.MaxUint64} {
				var actual bytes.Buffer
				_ = binary.Write(&actual, binary.LittleEndian, int64(expected))
				t.Run("test with value", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
				t.Run("test with pointer", func(t *testing.T) {
					expectedBytes, err := KeyToBytes(&expected)
					if err != nil {
						t.Fatalf("expected error for integer type but got '%v'", err)
					}
					if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
						t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
					}
				})
			}
		})

	})
	t.Run("test with floating-point type", func(t *testing.T) {
		t.Run("test with float32", func(t *testing.T) {
			var actual bytes.Buffer
			expected := float32(3.14)
			_ = binary.Write(&actual, binary.LittleEndian, expected)
			expectedBytes, err := KeyToBytes(expected)
			if err != nil {
				t.Fatalf("expected error for integer type but got '%v'", err)
			}
			if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
				t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
			}
		})
		t.Run("test with float64", func(t *testing.T) {
			var actual bytes.Buffer
			expected := float64(3.14)
			_ = binary.Write(&actual, binary.LittleEndian, expected)
			expectedBytes, err := KeyToBytes(expected)
			if err != nil {
				t.Fatalf("expected error for integer type but got '%v'", err)
			}
			if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
				t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
			}
		})

	})
	t.Run("test with string type", func(t *testing.T) {
		t.Run("test with string (empty)", func(t *testing.T) {
			var actual bytes.Buffer
			expected := words.EmptyString
			actual.WriteString(expected)
			expectedBytes, err := KeyToBytes(expected)
			if err != nil {
				t.Fatalf("expected error for integer type but got '%v'", err)
			}
			if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
				t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
			}
		})
		t.Run("test with string (non-empty)", func(t *testing.T) {
			var actual bytes.Buffer
			expected := "test string"
			actual.WriteString(expected)
			expectedBytes, err := KeyToBytes(expected)
			if err != nil {
				t.Fatalf("expected error for integer type but got '%v'", err)
			}
			if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
				t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
			}
		})
	})
	t.Run("test with rune type", func(t *testing.T) {
		t.Run("test with rune", func(t *testing.T) {
			var actual bytes.Buffer
			expected := 'a'
			_ = binary.Write(&actual, binary.LittleEndian, expected)
			expectedBytes, err := KeyToBytes(expected)
			if err != nil {
				t.Fatalf("expected error for integer type but got '%v'", err)
			}
			if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
				t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
			}
		})
		t.Run("test with rune", func(t *testing.T) {
			var actual bytes.Buffer
			expected := rune(0x00)
			_ = binary.Write(&actual, binary.LittleEndian, expected)
			expectedBytes, err := KeyToBytes(expected)
			if err != nil {
				t.Fatalf("expected error for integer type but got '%v'", err)
			}
			if bytes.Compare(actual.Bytes(), expectedBytes) != 0 {
				t.Fatalf("expected '%v' but got '%v'", expectedBytes, actual)
			}
		})
	})
	t.Run("test with struct", func(t *testing.T) {
		t.Run("test with struct type", func(t *testing.T) {
			type structType struct{}
			var lhs structType
			if _, err := KeyToBytes(lhs); err != nil {
				if err.Error() != errors.UnsupportedType {
					t.Fatalf("expected UnsupportedType error but got %v", err)
				}
			}
		})
		t.Run("test with struct pointer type", func(t *testing.T) {
			type structType struct{}
			var lhs = &(structType{})
			if _, err := KeyToBytes(lhs); err != nil {
				if err.Error() != errors.UnsupportedType {
					t.Fatalf("expected UnsupportedType error but got %v", err)
				}
			}
		})
	})

}
