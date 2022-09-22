package unique_occurrences

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	tableTests := []struct {
		arr    []int
		result bool
	}{
		{arr: []int{1, 2, 2, 1, 1, 3}, result: true},
		{arr: []int{1, 2}, result: false},
		{arr: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, result: true},
	}

	for _, tt := range tableTests {
		r := uniqueOccurrences(tt.arr)
		require.Equal(t, tt.result, r)
	}
}
