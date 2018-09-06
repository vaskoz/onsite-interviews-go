package sortedletters

import "testing"

var testcases = []struct {
	input    []string
	expected string
}{
	{[]string{"foo", "bar", "baz"}, "aabbfoorz"},
}

func TestSortLetters(t *testing.T) {
	for _, tc := range testcases {
		if result := SortLetters(tc.input); tc.expected != result {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkSortLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SortLetters(tc.input)
		}
	}
}
