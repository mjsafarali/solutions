package contains_duplicate_iii

import "math"

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j-i <= indexDiff && j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(valueDiff) {
				return true
			}
		}
	}

	return false
}
