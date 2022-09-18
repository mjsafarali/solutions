package palindrome

import "strings"

func isPalindrome(s string) bool {
	ns := strings.ReplaceAll(s, " ", "")
	ns = clean([]byte(ns))
	ns = strings.ToLower(ns)
	if reverseString(ns) == ns {
		return true
	}

	return false
}

func clean(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

func reverseString(word string) (reversed string) {
	for _, v := range word {
		reversed = string(v) + reversed
	}

	return
}
