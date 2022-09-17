package contains_duplicate_ii

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tableTests := []struct {
		nums                []int
		k                   int
		hasNearbyDuplicated bool
	}{
		{nums: []int{}, k: 10, hasNearbyDuplicated: false},
		{nums: []int{1}, k: 2, hasNearbyDuplicated: false},
		{nums: []int{1, 2, 3}, k: 3, hasNearbyDuplicated: false},
		{nums: []int{1, 2, 3, 1}, k: 3, hasNearbyDuplicated: true},
		{nums: []int{1, 0, 1, 1}, k: 1, hasNearbyDuplicated: true},
		{nums: []int{1, 2, 3, 1, 2, 3}, k: 2, hasNearbyDuplicated: false},
	}

	for _, tt := range tableTests {
		r := containsNearbyDuplicate(tt.nums, tt.k)
		require.Equal(t, tt.hasNearbyDuplicated, r)
	}
}
