package sortedletters

import "strings"

// SortLetters should return a string representing all the letters in sorted order.
func SortLetters(strs []string) string {
	letters := make([]int, 26) // only 'a'-'z'
	for _, str := range strs {
		for _, letter := range str {
			if letter >= 'a' && letter <= 'z' {
				letters[int(letter-'a')]++
			}
		}
	}
	var sb strings.Builder
	for i, count := range letters {
		if count == 0 {
			continue
		}
		sb.WriteString(strings.Repeat(string('a'+i), count))
	}
	return sb.String()
}
