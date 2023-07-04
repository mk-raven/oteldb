package logql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidLabel(t *testing.T) {
	tests := []struct {
		input    string
		allowDot bool
		wantErr  string
	}{
		{`a`, false, ""},
		{`A`, false, ""},
		{`_`, false, ""},
		{`a_b`, false, ""},
		{`a.b`, true, ""},

		{`a.`, false, "invalid label name character '.' at 1"},
		{`0`, false, "invalid label name character '0' at 0"},
		{``, false, "label name cannot be empty"},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			err := IsValidLabel(tt.input, tt.allowDot)
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}
