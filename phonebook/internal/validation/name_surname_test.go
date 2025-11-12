package validation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateNameSurname(t *testing.T) {
	test := []struct {
		name  string
		want  bool
		value string
	}{
		{
			name:  "valid ru",
			want:  true,
			value: "Хорьков",
		},
		{
			name:  "valid eng",
			want:  true,
			value: "Khorkov",
		},
		{
			name:  "invalid ru",
			want:  false,
			value: "хорьков",
		},
		{
			name:  "invalid eng",
			want:  false,
			value: "khorkov",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateNameSurname(tt.value)
			require.Equal(t, tt.want, got)
		})
	}
}
