package perfect_square

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	for i := 1; i <= int(float64(num)/2); i++ {
		if num == i*i {
			return true
		}
	}

	return false
}
