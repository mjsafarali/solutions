package palindrome

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPalindromic(t *testing.T) {
	tableTests := []struct {
		string   string
		expected bool
	}{
		{string: "A man, a plan, a canal: Panama", expected: true},
		{string: "race a car", expected: false},
		{string: " ", expected: true},
	}

	for _, tt := range tableTests {
		r := isPalindrome(tt.string)
		require.Equal(t, tt.expected, r)
	}
}
