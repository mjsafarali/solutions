package check_perfect_number

func checkPerfectNumber(num int) bool {
	var sum int
	for i := 1; i <= num/2; i++ {
		dv := float64(num) / float64(i)
		if float64(int(dv)) == dv {
			sum = sum + i
		}
	}

	if sum == num {
		return true
	}

	return false
}
