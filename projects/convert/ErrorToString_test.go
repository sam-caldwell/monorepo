package convert

import (
	"errors"
	"testing"

	"github.com/sam-caldwell/go/v2/projects/misc/words"
)

func TestErrorToString(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "With error",
			err:  errors.New("test error"),
			want: "test error",
		},
		{
			name: "Without error",
			err:  nil,
			want: words.EmptyString,
		},
	}

	for _, tt := range tests {
		if got := ErrorToString(tt.err); tt.want != got {
			t.Fatalf("test failed: %s\n"+
				"\t got:%v\n"+
				"\twant:%v\n",
				tt.name, got, tt.want)
		}
	}
}
