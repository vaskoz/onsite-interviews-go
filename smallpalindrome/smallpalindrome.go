package smallpalindrome

// SmallestPalindrome returns the smallest possible palindrome by adding to the prefix.
func SmallestPalindrome(str string) string {
	var result string
	for i := len(str); i >= 0; i-- {
		substr := str[i:] // from i to the end
		if result = reverse(substr) + str; isPalindrome(result) {
			break
		}
	}
	return result
}

func reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
