package judge_square_sum

import "math"

func judgeSquareSum(c int) bool {
	result := false

	for i := 0; c-(i*i) >= 0; i++ {
		a2 := c - (i * i)
		val := math.Sqrt(float64(a2))
		if val == float64(int(val)) {
			result = true
			break
		}
	}

	return result
}
