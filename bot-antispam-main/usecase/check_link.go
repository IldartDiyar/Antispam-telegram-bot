package usecase

import "regexp"

func CheckLink(msg string) (string, bool) {
	linkPattern := `((http|https)\:\/\/)?[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?`
	linkRegex := regexp.MustCompile(linkPattern)
	var foundLink string
	if linkRegex.MatchString(msg) {
		foundLink = linkRegex.FindString(msg)
		return foundLink, true
	}
	return "", false
}
