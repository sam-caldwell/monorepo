package pair

import (
	"testing"
)

func TestCompareKey(t *testing.T) {
	t.Run("Test with boolean values", func(t *testing.T) {
		t.Run("test with value", func(t *testing.T) {
			t.Run("test lhs==rhs", func(t *testing.T) {
				expect := CompareKeyEqual
				lhs, rhs := true, true
				if v := CompareKey[bool](lhs, rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs<rhs", func(t *testing.T) {
				expect := CompareKeyLess
				lhs, rhs := false, true
				if v := CompareKey[bool](lhs, rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs>rhs", func(t *testing.T) {
				expect := CompareKeyGreater
				lhs, rhs := true, false
				if v := CompareKey[bool](lhs, rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
		})
		t.Run("test with pointer", func(t *testing.T) {
			t.Run("test lhs==rhs", func(t *testing.T) {
				expect := CompareKeyEqual
				lhs, rhs := true, true
				if v := CompareKey[*bool](&lhs, &rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs<rhs", func(t *testing.T) {
				expect := CompareKeyLess
				lhs, rhs := false, true
				if v := CompareKey[*bool](&lhs, &rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs>rhs", func(t *testing.T) {
				expect := CompareKeyGreater
				lhs, rhs := true, false
				if v := CompareKey[*bool](&lhs, &rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
		})
	})
	t.Run("Test with numeric (signed integer) values", func(t *testing.T) {
		t.Run("test with int", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := -123, -123
					if v := CompareKey[int](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := -123, +123
					if v := CompareKey[int](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => 1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := 123, -123
					if v := CompareKey[int](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := -123, -123
					if v := CompareKey[*int](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := -123, +123
					if v := CompareKey[*int](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => 1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := 123, -123
					if v := CompareKey[*int](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with int8", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int8(-123), int8(-123)
					if v := CompareKey[int8](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => 1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int8(-123), int8(123)
					if v := CompareKey[int8](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => 1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int8(123), int8(-123)
					if v := CompareKey[int8](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int8(-123), int8(-123)
					if v := CompareKey[*int8](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => 1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int8(-123), int8(123)
					if v := CompareKey[*int8](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => 1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int8(123), int8(-123)
					if v := CompareKey[*int8](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with int16", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int16(-123), int16(-123)
					if v := CompareKey[int16](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int16(-123), int16(123)
					if v := CompareKey[int16](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int16(123), int16(-123)
					if v := CompareKey[int16](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int16(-123), int16(-123)
					if v := CompareKey[*int16](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int16(-123), int16(123)
					if v := CompareKey[*int16](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int16(123), int16(-123)
					if v := CompareKey[*int16](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with int32", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int32(-123), int32(-123)
					if v := CompareKey[int32](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int32(-123), int32(123)
					if v := CompareKey[int32](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int32(123), int32(-123)
					if v := CompareKey[int32](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int32(-123), int32(-123)
					if v := CompareKey[*int32](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int32(-123), int32(123)
					if v := CompareKey[*int32](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int32(123), int32(-123)
					if v := CompareKey[*int32](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with int64", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int64(-123), int64(-123)
					if v := CompareKey[int64](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int64(-123), int64(123)
					if v := CompareKey[int64](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int64(123), int64(-123)
					if v := CompareKey[int64](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := int64(-123), int64(-123)
					if v := CompareKey[*int64](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int64(-123), int64(123)
					if v := CompareKey[*int64](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := int64(123), int64(-123)
					if v := CompareKey[*int64](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
	})
	t.Run("Test with numeric (unsigned integer) values", func(t *testing.T) {
		t.Run("test with uint", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint(32), uint(32)
					if v := CompareKey[uint](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint(32), uint(64)
					if v := CompareKey[uint](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint(64), uint(32)
					if v := CompareKey[uint](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint(32), uint(32)
					if v := CompareKey[*uint](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint(32), uint(64)
					if v := CompareKey[*uint](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint(64), uint(32)
					if v := CompareKey[*uint](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with uint8", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint8(32), uint8(32)
					if v := CompareKey[uint8](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint8(32), uint8(64)
					if v := CompareKey[uint8](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint8(64), uint8(32)
					if v := CompareKey[uint8](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint8(32), uint8(32)
					if v := CompareKey[*uint8](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint8(32), uint8(64)
					if v := CompareKey[*uint8](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint8(64), uint8(32)
					if v := CompareKey[*uint8](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with uint16", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint16(32), uint16(32)
					if v := CompareKey[uint16](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint16(32), uint16(64)
					if v := CompareKey[uint16](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint16(64), uint16(32)
					if v := CompareKey[uint16](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint16(32), uint16(32)
					if v := CompareKey[*uint16](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint16(32), uint16(64)
					if v := CompareKey[*uint16](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint16(64), uint16(32)
					if v := CompareKey[*uint16](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with uint32", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint32(32), uint32(32)
					if v := CompareKey[uint32](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint32(32), uint32(64)
					if v := CompareKey[uint32](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint32(64), uint32(32)
					if v := CompareKey[uint32](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint32(32), uint32(32)
					if v := CompareKey[*uint32](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint32(32), uint32(64)
					if v := CompareKey[*uint32](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint32(64), uint32(32)
					if v := CompareKey[*uint32](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
		t.Run("test with uint64", func(t *testing.T) {
			t.Run("test with value", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint64(32), uint64(32)
					if v := CompareKey[uint64](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint64(32), uint64(64)
					if v := CompareKey[uint64](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint64(64), uint64(32)
					if v := CompareKey[uint64](lhs, rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
			t.Run("test with pointer", func(t *testing.T) {
				t.Run("test lhs==rhs", func(t *testing.T) {
					expect := CompareKeyEqual
					lhs, rhs := uint64(32), uint64(32)
					if v := CompareKey[*uint64](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs<rhs", func(t *testing.T) {
					expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint64(32), uint64(64)
					if v := CompareKey[*uint64](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
				t.Run("test lhs>rhs", func(t *testing.T) {
					expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
					lhs, rhs := uint64(64), uint64(32)
					if v := CompareKey[*uint64](&lhs, &rhs); v != expect {
						t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
					}
				})
			})
		})
	})
	t.Run("Test with string", func(t *testing.T) {
		t.Run("test with value", func(t *testing.T) {
			t.Run("test lhs==rhs", func(t *testing.T) {
				expect := CompareKeyEqual
				lhs, rhs := "testData", "testData"
				if v := CompareKey[string](lhs, rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs<rhs", func(t *testing.T) {
				expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
				lhs, rhs := "testDataA", "testDataB"
				if v := CompareKey[string](lhs, rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs>rhs", func(t *testing.T) {
				expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
				lhs, rhs := "testDataB", "testDataA"
				if v := CompareKey[string](lhs, rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
		})
		t.Run("test with pointer", func(t *testing.T) {
			t.Run("test lhs==rhs", func(t *testing.T) {
				expect := CompareKeyEqual
				lhs, rhs := "testData", "testData"
				if v := CompareKey[*string](&lhs, &rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs<rhs", func(t *testing.T) {
				expect := CompareKeyLess //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
				lhs, rhs := "testDataA", "testDataB"
				if v := CompareKey[*string](&lhs, &rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
			t.Run("test lhs>rhs", func(t *testing.T) {
				expect := CompareKeyGreater //lhs<rhs => -1 but because we compare using bytes, the signing bit inverts order
				lhs, rhs := "testDataB", "testDataA"
				if v := CompareKey[*string](&lhs, &rhs); v != expect {
					t.Fatalf("CompareKey(lhs,rhs)=%d, want %d", v, expect)
				}
			})
		})
	})
}
