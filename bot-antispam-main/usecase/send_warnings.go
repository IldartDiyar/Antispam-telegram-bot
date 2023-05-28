package usecase

import "fmt"

func PrintMap(maps map[string]int) string {
	var message string
	for key, value := range maps {
		message += fmt.Sprintf("%s: %d\n", key, value)
	}
	return message
}
