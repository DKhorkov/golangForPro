package validation

import "regexp"

var (
	engNamePattern = regexp.MustCompile(`^[A-Z][a-z]*$`)
	ruNamePattern  = regexp.MustCompile(`^[А-Я][а-я]*$`)
)

func ValidateNameSurname(value string) bool {
	b := []byte(value)
	return engNamePattern.Match(b) || ruNamePattern.Match(b)
}
