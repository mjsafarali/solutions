package judge_square_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPalindromic(t *testing.T) {
	tableTests := []struct {
		c        int
		expected bool
	}{
		{c: 5, expected: true},
		{c: 3, expected: false},
		{c: 4, expected: true},
		{c: 0, expected: true},
	}

	for _, tt := range tableTests {
		r := judgeSquareSum(tt.c)
		require.Equal(t, tt.expected, r)
	}
}
