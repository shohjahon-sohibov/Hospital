package helper

import (
	"fmt"
	"regexp"
)

func ExtractPaths(text string) []string {
	pattern := `<path[^>]*\/>`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(text, -1)
	return matches
}

func ExtractID(input string) (string, error) {
	re := regexp.MustCompile(`id="([^"]+)"`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 2 {
		return "", fmt.Errorf("ID not found")
	}
	return matches[1], nil
}
