package common

import "regexp"

// Remove deletes an element from a slice. It does not maintain order
func Remove(s []interface{}, i int) []interface{} {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

func ValidateHTTP(test string) bool {
	tester := regexp.MustCompile(`(?m)(https)|(http)`)
	return tester.Match([]byte(test))
}
