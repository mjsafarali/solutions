package palindromic

func firstPalindrome(words []string) string {
	var firstPalindrome string
	for _, word := range words {
		if reverseString(word) == word {
			firstPalindrome = word
			break
		}
	}

	return firstPalindrome
}

func reverseString(word string) (reversed string) {
	for _, v := range word {
		reversed = string(v) + reversed
	}

	return
}
