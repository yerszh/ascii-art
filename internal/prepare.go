package asciiArt

import (
	"os"
	"regexp"
	"strings"
)

func onlyNewLines(input []string) bool {
	for _, ch := range input {
		if ch != "" {
			return false
		}
	}
	return true
}

func PrepareForOutput(banner string, text string) ([]string, []string) {
	dataBytes, _ := os.ReadFile("assets/" + banner + ".txt")
	standardAscii := []string{}

	if banner == "thinkertoy" {
		standardAscii = strings.Split(string(dataBytes), "\r\n")
	} else {
		standardAscii = strings.Split(string(dataBytes), "\n")
	}

	inputWords := text
	splittedWords := regexp.MustCompile(`\n|\\n`).Split(inputWords, -1)

	if onlyNewLines(splittedWords) {
		splittedWords = splittedWords[1:]
	}

	return splittedWords, standardAscii
}