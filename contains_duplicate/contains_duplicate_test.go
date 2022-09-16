package contains_duplicate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tableTests := []struct {
		nums          []int
		hasDuplicated bool
	}{
		{nums: []int{}, hasDuplicated: false},
		{nums: []int{1}, hasDuplicated: false},
		{nums: []int{1, 2, 3}, hasDuplicated: false},
		{nums: []int{1, 2, 3, 2}, hasDuplicated: true},
		{nums: []int{1, 2, 3, 2, 3, 3, 3}, hasDuplicated: true},
	}

	for _, tt := range tableTests {
		r := containsDuplicate(tt.nums)
		require.Equal(t, tt.hasDuplicated, r)
	}
}
