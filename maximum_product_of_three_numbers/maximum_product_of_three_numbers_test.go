package maximum_product_of_three_numbers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaximumProduct(t *testing.T) {
	tableTests := []struct {
		nums   []int
		result int
	}{
		{nums: []int{1, 2, 3}, result: 6},
		{nums: []int{1, 2, 3, 4}, result: 24},
		{nums: []int{-1, -2, -3}, result: -6},
	}

	for _, tt := range tableTests {
		r := maximumProduct(tt.nums)
		require.Equal(t, tt.result, r)
	}
}
