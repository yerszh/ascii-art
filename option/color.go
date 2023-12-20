package asciiArt

import (
	"fmt"
	"os"
	"strings"

	internal "ascii-art-output/internal"
)

func Atoi(text string) int {
	result := 0
	for _, char := range text {
		if char >= '0' && char <= '9' {
			number := int(char - '0')
			result = result*10 + number
		}
	}
	return result
}

func FindRGB(color string) (int, int, int) {
	arrayOfRGBSystem := []string{}
	r, g, b := 0, 0, 0

	if color[len(color)-1] != ')' {
		r = Atoi(os.Args[1])
		g = Atoi(os.Args[2])
		b = Atoi(os.Args[3])
	} else {
		arrayOfRGBSystem = strings.Split(color, ",")
		if len(arrayOfRGBSystem) == 3 {
			r = Atoi(arrayOfRGBSystem[0])
			g = Atoi(arrayOfRGBSystem[1])
			b = Atoi(arrayOfRGBSystem[2])
		}
	}
	return r, g, b
}

func PickColor(color string) string {
	switch color {
	case "black":
		return "\u001b[30m"
	case "red":
		return "\u001b[31m"
	case "green":
		return "\u001b[32m"
	case "yellow":
		return "\u001b[33m"
	case "blue":
		return "\u001b[34m"
	case "magenta":
		return "\u001b[35m"
	case "cyan":
		return "\u001b[36m"
	case "white":
		return "\u001b[37m"
	}
	return "Failed"
}

func rgbToAnsi(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}


func Color() {
	color := strings.ToLower(internal.FindFile(os.Args[1]))

	if color == "failed" {
		fmt.Println("Missing argument\nTry like this  --> \" --color=rgb(48, 255, 229)\"")
		os.Exit(1)
	}

	var ansiCode string
	var r, g, b int
	var letterToBeColored string
	var text string
	var banner string

	if color[len(color)-1] == ',' {
		r, g, b = FindRGB(color)
		ansiCode = rgbToAnsi(r, g, b)

		switch len(os.Args) {
		case 7:
			letterToBeColored = os.Args[4]
			text = os.Args[5]
			banner = os.Args[6]
		case 6:
			if internal.CheckIsBanner(os.Args[len(os.Args)-1]) {
				text = os.Args[4]
				banner = os.Args[5]
			} else {
				letterToBeColored = os.Args[4]
				text = os.Args[5]
				banner = "standard"
			}
		case 5:
			text = os.Args[4]
			banner = "standard"
		}
	} else {
		if color[len(color)-1] == ')' {
			r, g, b = FindRGB(color)
			ansiCode = rgbToAnsi(r, g, b)
		} else {
			ansiCode = PickColor(color)
			if ansiCode == "Failed" {
				fmt.Println("Wrong color\nChoose one of these colors --> black, red, gree, yellow, blue, magenta, cyan, white")
				os.Exit(1)
			}
		}

		switch len(os.Args) {
		case 5:
			letterToBeColored = os.Args[2]
			text = os.Args[3]
			banner = os.Args[4]
		case 4:
			if os.Args[len(os.Args)-1] == "standard" || os.Args[len(os.Args)-1] == "shadow" || os.Args[len(os.Args)-1] == "thinkertoy" {
				text = os.Args[2]
				banner = os.Args[3]
			} else {
				letterToBeColored = os.Args[2]
				text = os.Args[3]
				banner = "standard"
			}
		case 3:
			text = os.Args[2]
			banner = "standard"
		}
	}

	resetCode := "\x1b[0m"
	text, notValidText := internal.CheckIsAscii(text)

	if notValidText != "" {
		fmt.Println(notValidText)
		os.Exit(1)
	}

	if !internal.CheckIsBanner(banner) {
		fmt.Println("Wrong number of arguments\nUsage: \"go run . --color=rgb(48, 255, 229) smthng something thinkertoy\"")
		os.Exit(1)
	}

	if internal.CheckForChangeFile("assets/"+banner+".txt", banner) {
		splitedWords, standardAscii := internal.PrepareForOutput(banner, text)

		for _, word := range splitedWords {
			if word == "" {
				fmt.Println()
				continue
			}

			for index := 1; index <= 8; index++ {
				for _, ch := range word {
					colored := false
					if len(letterToBeColored) != 0 {
						for _, letter := range letterToBeColored {
							if ch == letter {
								colored = true
								break
							}
						}
					} else {
						colored = true
					}

					if colored {
						fmt.Print(ansiCode, standardAscii[int((ch-32)*9)+index])
					} else {
						fmt.Print(resetCode, standardAscii[int((ch-32)*9)+index])
					}
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("The file has been changed , the program will close")
		os.Exit(1)
	}
}