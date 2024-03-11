package gox

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// SplitStringsToMap runs the [strings.Cut] function on a slice of strings and
// for each entry it sets a map entry. Everything before the [sep] is placed as
// the key, everything after is the value.
//
// If any of the entries only return 1 string, then an empty value is set for
// the key placement.
func SplitStringsToMap(strs []string, sep string) (ret map[string]string) {
	ret = make(map[string]string)

	var before, after string
	var ok bool
	for _, str := range strs {
		if before, after, ok = strings.Cut(str, sep); ok {
			ret[before] = after
		} else {
			ret[str] = ""
		}
	}
	return
}

// UniquifyString makes a string "unique"-ish for conflict checking by normalizing
// the text to ASCII lower-case.
func UniquifyString(str string) (string, error) {
	result, _, err := transform.String(transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn))), str)
	return strings.ToLower(result), err
}

// AfterRune returns everything after the last instance of the given substring. If the
// rune does not exist, the string is returned unchanged.
func After(str, sub string) string {
	if ind := strings.LastIndex(str, sub); ind > 0 {
		return str[ind+1:]
	}
	return str
}

// SplitStringByRune splits up the string by searching for the given rune. It
// does not consume the rune in the parts provided. Remaining tail after the
// last rune is included.
func SplitStringByRune(str string, run rune) []string {
	parts := make([]string, 0)
	lastInd := 0
	for i, r := range str {
		if r == run {
			parts = append(parts, str[lastInd:i])
			lastInd = i + 1
		}
	}
	if lastInd < len(str)-1 {
		parts = append(parts, str[lastInd:])
	}
	return parts
}
