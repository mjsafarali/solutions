package unique_occurrences

func uniqueOccurrences(arr []int) bool {
	intWithCount := make(map[int]int) //map[inputNumber]Occurrences
	dCheck := make(map[int]bool)

	for _, num := range arr {
		intWithCount[num] = intWithCount[num] + 1
	}

	for _, oc := range intWithCount {
		if dCheck[oc] == true {
			return false
		}
		dCheck[oc] = true
	}

	return true
}
