package q5

import "strings"

func reverse(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	specialChars := []string{":", ",", " "}

	for _, char := range specialChars {
		s = strings.ReplaceAll(s, char, "")
	}

	if reverse(s) == s {
		return true
	}

	return false
}
