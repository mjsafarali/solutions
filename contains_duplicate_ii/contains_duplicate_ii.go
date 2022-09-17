package contains_duplicate_ii

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	a := map[int]int{}
	for i, num := range nums {
		if j, found := a[num]; found && math.Abs(float64(i-j)) <= float64(k) {
			return true
		}
		a[num] = i
	}

	return false
}
