package smallpalindrome

import "testing"

var testcases = []struct {
	in, expected string
}{
	{"abcd", "dcbabcd"},
	{"lol", "lol"},
	{"vasko", "oksavasko"},
	{"car", "racar"},
}

func TestSmallestPalindrome(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SmallestPalindrome(tc.in); result != tc.expected {
			t.Errorf("For input %v expected %v but got %v", tc.in, tc.expected, result)
		}
	}
}

func BenchmarkSmallestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestPalindrome(tc.in)
		}
	}
}
