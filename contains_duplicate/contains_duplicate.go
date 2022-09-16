package contains_duplicate

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	a := make(map[int]bool)

	for _, num := range nums {
		if a[num] == true {
			return true
		}

		a[num] = true
	}

	return false
}
