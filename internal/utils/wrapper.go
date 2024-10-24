package utils

import "strings"

func WrapText(text string, lineWidth int) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	var result string
	var line string
	for _, word := range words {
		if len(line)+len(word)+1 > lineWidth {
			result += line + "\n"
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}
	if line != "" {
		result += line
	}
	return result
}
