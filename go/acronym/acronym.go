package acronym

import "strings"
import "unicode"

const testVersion = 3

func Abbreviate(s string) string {
	acronym, take := []rune{}, true
	for _, r := range s {
		if take && isspace(r) == false {
			acronym = append(acronym, r)
		}
		take = isspace(r)
	}
	return strings.ToUpper(string(acronym))
}

func isspace(r rune) bool {
	if unicode.IsSpace(r) || r == '-' {
		return true
	}
	return false
}
