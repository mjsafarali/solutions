package check_perfect_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckPerfectNumber(t *testing.T) {
	tableTests := []struct {
		num    int
		result bool
	}{
		{num: 28, result: true},
		{num: 7, result: false},
	}

	for _, tt := range tableTests {
		r := checkPerfectNumber(tt.num)
		require.Equal(t, tt.result, r)
	}
}
