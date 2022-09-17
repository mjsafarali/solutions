package palindromic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPalindromic(t *testing.T) {
	tableTests := []struct {
		words  []string
		result string
	}{
		{words: []string{"abc", "car", "ada", "racecar", "cool"}, result: "ada"},
		{words: []string{"notapalindrome", "racecar"}, result: "racecar"},
		{words: []string{"def", "ghi"}, result: ""},
	}

	for _, tt := range tableTests {
		r := firstPalindrome(tt.words)
		require.Equal(t, tt.result, r)
	}
}
