package contains_duplicate_iii

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tableTests := []struct {
		nums                      []int
		indexDiff                 int
		valueDiff                 int
		hasNearbyAlmostDuplicated bool
	}{
		{nums: []int{}, indexDiff: 10, valueDiff: 10, hasNearbyAlmostDuplicated: false},
		{nums: []int{1}, indexDiff: 10, valueDiff: 20, hasNearbyAlmostDuplicated: false},
		{nums: []int{1, 2, 3, 1}, indexDiff: 3, valueDiff: 0, hasNearbyAlmostDuplicated: true},
		{nums: []int{1, 5, 9, 1, 5, 9}, indexDiff: 2, valueDiff: 3, hasNearbyAlmostDuplicated: false},
	}

	for _, tt := range tableTests {
		r := containsNearbyAlmostDuplicate(tt.nums, tt.indexDiff, tt.valueDiff)
		require.Equal(t, tt.hasNearbyAlmostDuplicated, r)
	}
}
