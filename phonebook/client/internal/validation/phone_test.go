package validation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatePhone(t *testing.T) {
	test := []struct {
		name  string
		want  bool
		phone string
	}{
		{
			name:  "valid",
			want:  true,
			phone: "+7 (911) 258-01-62",
		},
		{
			name:  "invalid",
			want:  false,
			phone: "123",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePhone(tt.phone)
			require.Equal(t, tt.want, got)
		})
	}
}
