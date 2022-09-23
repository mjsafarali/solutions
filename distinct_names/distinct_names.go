package distinct_names

func distinctNames(ideas []string) int64 {
	var result int64

	for _, ideaA := range ideas {
		for _, ideaB := range ideas {
			if ideaA == ideaB {
				continue
			}

			nStr1, nStr2 := swapFirstLetter(ideaA, ideaB)
			if stringInSlice(nStr1, ideas) == false && stringInSlice(nStr2, ideas) == false {
				result++
			}
		}
	}

	return result
}

func swapFirstLetter(str1 string, str2 string) (nstr1 string, nstr2 string) {
	fcStr1, str1 := trimLeftCharAndReturnFc(str1)
	fcStr2, str2 := trimLeftCharAndReturnFc(str2)

	return fcStr2 + str1, fcStr1 + str2
}

func trimLeftCharAndReturnFc(s string) (fcStr string, newStr string) {
	fcStr = string(s[0])

	if len(s) > 1 {
		return fcStr, s[1:]
	}

	return fcStr, ""
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
