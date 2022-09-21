package maximum_product_of_three_numbers

import (
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	return max(nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3], nums[0]*nums[1]*nums[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
