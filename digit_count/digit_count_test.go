package digit_count

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDigitCount(t *testing.T) {
	testsTable := []struct {
		num    string
		result bool
	}{
		{num: "1210", result: true},
		{num: "030", result: false},
	}
	for _, tt := range testsTable {
		a := digitCount(tt.num)
		require.Equal(t, tt.result, a)
	}
}
