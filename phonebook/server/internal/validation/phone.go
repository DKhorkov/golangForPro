package validation

import "regexp"

var (
	phonePattern = regexp.MustCompile(`^\+7 \(\d{3}\) \d{3}-\d{2}-\d{2}$`)
)

func ValidatePhone(phone string) bool {
	return phonePattern.Match([]byte(phone))
}
