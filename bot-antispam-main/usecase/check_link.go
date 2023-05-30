package usecase

import (
	"regexp"
)

func CheckLink(msg string) (string, bool) {
	pattern := `((http|https):\/\/[^\s]+)`

	regex := regexp.MustCompile(pattern)

	match := regex.FindString(msg)

	if match != "" {
		return match, true
	}

	return "", false
}
