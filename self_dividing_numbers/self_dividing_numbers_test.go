package self_dividing_numbers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	testsTable := []struct {
		left   int
		right  int
		result []int
	}{
		{left: 1, right: 22, result: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{left: 47, right: 85, result: []int{48, 55, 66, 77}},
	}

	for _, tt := range testsTable {
		a := selfDividingNumbers(tt.left, tt.right)
		require.Equal(t, tt.result, a)
	}
}
