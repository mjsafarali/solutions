package word_pattern

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	mw := map[string]string{} //map based on words: dog => a
	mp := map[string]string{} //map based on pattern: a => dog

	for i, word := range words {
		patternI := string(pattern[i])

		fPattern, fExists := mw[word]  //fPattern = found pattern bases on word
		fWord, cExists := mp[patternI] //fWord = found word based on pattern

		if fExists && cExists {
			if fPattern == patternI && fWord == word {
				words[i] = patternI
			}
		} else if !fExists && !cExists {
			mw[word] = patternI
			mp[patternI] = word
			words[i] = patternI
		}
	}

	return strings.Join(words, "") == pattern
}
