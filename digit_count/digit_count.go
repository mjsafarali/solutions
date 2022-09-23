package digit_count

import (
	"strconv"
)

func digitCount(num string) bool {
	result := true

	for i := 0; i < len(num); i++ {
		expectedOccur, _ := strconv.Atoi(string(num[i]))
		actualOccur := 0

		for j := 0; j < len(num); j++ {
			item, _ := strconv.Atoi(string(num[j]))
			if item == i {
				actualOccur++
			}
		}

		if expectedOccur != actualOccur {
			result = false
			break
		}
	}

	return result
}
