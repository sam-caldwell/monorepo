package types

import (
	"fmt"
	"testing"
)

func TestArgTypes_Typecheck(t *testing.T) {
	type expected struct {
		t ArgTypes // type associated with v for type check
		v any      // test value we will type check
		f bool     // expected failure
	}
	var err error
	var data = []expected{
		{t: Boolean, v: true, f: false},
		{t: Boolean, v: false, f: false},
		{t: Boolean, v: -1.0, f: true},
		{t: Boolean, v: 0.0, f: true},
		{t: Boolean, v: 1.0, f: true},
		{t: Boolean, v: -3, f: true},
		{t: Boolean, v: 0, f: true},
		{t: Boolean, v: 3, f: true},
		{t: Boolean, v: "true", f: true},
		{t: Boolean, v: "false", f: true},

		{t: Flag, v: true, f: false},
		{t: Flag, v: false, f: false},
		{t: Flag, v: -1.0, f: true},
		{t: Flag, v: 0.0, f: true},
		{t: Flag, v: 1.0, f: true},
		{t: Flag, v: -3, f: true},
		{t: Flag, v: 0, f: true},
		{t: Flag, v: 3, f: true},
		{t: Flag, v: "true", f: true},
		{t: Flag, v: "false", f: true},

		{t: Float, v: true, f: true},
		{t: Float, v: false, f: true},
		{t: Float, v: -1.0, f: false},
		{t: Float, v: 0.0, f: false},
		{t: Float, v: 1.0, f: false},
		{t: Float, v: -3, f: true},
		{t: Float, v: 0, f: true},
		{t: Float, v: 3, f: true},
		{t: Float, v: "true", f: true},
		{t: Float, v: "false", f: true},

		{t: Integer, v: true, f: true},
		{t: Integer, v: false, f: true},
		{t: Integer, v: -1.0, f: true},
		{t: Integer, v: 0.0, f: true},
		{t: Integer, v: 1.0, f: true},
		{t: Integer, v: -3, f: false},
		{t: Integer, v: 0, f: false},
		{t: Integer, v: 3, f: false},
		{t: Integer, v: "-3", f: true},
		{t: Integer, v: "0", f: true},
		{t: Integer, v: "3", f: true},
		{t: Integer, v: "test", f: true},

		{t: String, v: true, f: true},
		{t: String, v: false, f: true},
		{t: String, v: -1.0, f: true},
		{t: String, v: 0.0, f: true},
		{t: String, v: 1.0, f: true},
		{t: String, v: -3, f: true},
		{t: String, v: 0, f: true},
		{t: String, v: 3, f: true},
		{t: String, v: "-3", f: false},
		{t: String, v: "0", f: false},
		{t: String, v: "3", f: false},
		{t: Integer, v: "test", f: true},
		{t: Integer, v: "testMe", f: true},
	}

	for _, row := range data {
		func(o ArgTypes, v any, expectFail bool) {
			if err = o.Typecheck(v); err != nil {
				if expectFail {
					// we expect a failure, verify the error is ok
					func(arg ArgTypes, actualErr error, value any) {
						expectedErr := fmt.Errorf(eMsgTypeCheck, arg.String())
						if expectedErr.Error() != actualErr.Error() {
							t.Fatalf("'%s' '%s'", actualErr, expectedErr)
						}
					}(o, err, v)
				} else {
					t.Fatalf("Unexpected fail '%s' '%v' %v", o.String(), v, err)
				}
			} else {
				if expectFail {
					t.Fatalf("Expected fail not realized '%s' '%v'", o.String(), v)
				}
			}
		}(row.t, row.v, row.f)
	}
}
