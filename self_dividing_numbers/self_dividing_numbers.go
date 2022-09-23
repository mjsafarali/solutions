package self_dividing_numbers

import (
	"strconv"
	"strings"
)

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		numInString := strconv.Itoa(i)
		if strings.Contains(numInString, "0") {
			continue
		}

		var x []int
		shouldPass := false
		for _, c := range numInString {
			item, _ := strconv.Atoi(string(c))
			r := float64(i) / float64(item)
			if float64(int(r)) != r {
				shouldPass = true
			}
			x = append(x, item)
		}

		if shouldPass != false {
			continue
		}

		res = append(res, i)
	}

	return res
}
