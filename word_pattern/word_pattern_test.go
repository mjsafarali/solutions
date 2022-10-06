package word_pattern

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWordPattern(t *testing.T) {
	tableTests := []struct {
		pattern string
		words   string
		result  bool
	}{
		{pattern: "abba", words: "dog cat cat dog", result: true},
		{pattern: "abba", words: "dog cat cat fish", result: false},
		{pattern: "aaaa", words: "dog cat cat dog", result: false},
		{pattern: "aaa", words: "dog cat cat dog", result: false},
	}

	for _, tt := range tableTests {
		r := wordPattern(tt.pattern, tt.words)
		require.Equal(t, tt.result, r)
	}
}
