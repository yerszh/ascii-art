package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func formatWords(inputWords string) []string {
	words := regexp.MustCompile(`\n|\\n`).Split(inputWords, -1)
	onlyNewLines := true
	for _, word := range words {
		if word != "" {
			onlyNewLines = false
			break
		}
	}
	if onlyNewLines {
		return words[1:]
	}
	return words
}

func printAscii(inputWords, banner string) error {
	words := formatWords(inputWords)
	for _, word := range words {
		for _, r := range word {
			if r < 32 || 126 < r {
				return fmt.Errorf("error! use only latin letters")
			}
		}
	}

	standardAscii, _ := os.ReadFile("assets/" + banner + ".txt")
	toSplit := map[string]string{
		"standard":   "\n",
		"shadow":     "\n",
		"thinkertoy": "\r\n",
	}
	bannerData := strings.Split(string(standardAscii), toSplit[banner])

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		for index := 1; index <= 8; index++ {
			for _, ch := range word {
				fmt.Print(bannerData[int((ch-32)*9)+index])
			}
			fmt.Println()
		}
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		return
	}
	if len(os.Args[1]) > 30 {
		fmt.Println(os.Args[1])
		fmt.Println(`error! maximum number of characters is 30`)
		return
	}
	banner := "standard"
	if len(os.Args) == 3 {
		switch {
		case regexp.MustCompile(`(?i)\bstandard\b`).MatchString(os.Args[2]):
			banner = "standard"
		case regexp.MustCompile(`(?i)\bshadow\b`).MatchString(os.Args[2]):
			banner = "shadow"
		case regexp.MustCompile(`(?i)\bthinkertoy\b`).MatchString(os.Args[2]):
			banner = "thinkertoy"
		}
	}
	if len(os.Args) > 3 {
		fmt.Println(`error! invalid arguments, try: go run main.go "string" [banner]`)
		return
	}
	err := printAscii(os.Args[1], banner)
	if err != nil {
		fmt.Println(err)
		return
	}
}
