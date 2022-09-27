package perfect_square

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	testTable := []struct {
		num    int
		result bool
	}{
		{num: 16, result: true},
		{num: 14, result: false},
	}

	for _, tt := range testTable {
		r := isPerfectSquare(tt.num)
		require.Equal(t, tt.result, r)
	}
}
