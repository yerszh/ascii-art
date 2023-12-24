package asciiArt

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	internal "ascii-art-output/internal"
)

func SpaceCount(text string) int {
	counterForSpaceInText := 0
	for i := 0; i < len(text); i++ {
		if i != 0 && i <= len(text)-2 {
			if text[i-1] != ' ' && text[i] == ' ' {
				counterForSpaceInText++
			}
		}
	}
	return counterForSpaceInText
}

func CreateSpaces(lineFormArrayOfOutput string, terminalWidth int) string {
	spaces := ""
	for terminalWidth-len(lineFormArrayOfOutput) > 0 {
		spaces += " "
		terminalWidth--
	}
	return spaces
}

func GetTermWidth() int {
	out, er1 := exec.Command("tput", "cols").Output()
	out1 := strings.TrimSpace(string(out))
	num, er2 := strconv.Atoi(out1)
	if er1 != nil {
		os.Exit(0)
	}
	if er2 != nil {
		os.Exit(0)
	}
	return num
}


func Aling(align string, output string, text string) {
	arrayOfOutputWrong := strings.Split(output, "\n")
	arrayOfOutput := arrayOfOutputWrong[0 : len(arrayOfOutputWrong)-1]
	terminalWidth := GetTermWidth()
	countSpace := CreateSpaces(arrayOfOutput[0], terminalWidth)

	if align == "justify" && SpaceCount(text) == 0 {
		align = "center"
	}

	if terminalWidth < len(arrayOfOutput[0])+len(countSpace) {
		fmt.Println("Terminal width is too small")
		os.Exit(0)
	}

	switch align {
	case "left":
		fmt.Print(output)
	case "right":
		for _, lineFormArrayOfOutput := range arrayOfOutput {
			spacesLeft := CreateSpaces(lineFormArrayOfOutput, terminalWidth)
			fmt.Print(spacesLeft, lineFormArrayOfOutput)
		}
	case "center":
		for _, lineFormArrayOfOutput := range arrayOfOutput {
			spaces := CreateSpaces(lineFormArrayOfOutput, terminalWidth)
			spacesLeft := spaces[0 : len(spaces)/2]
			spacesRight := spaces[0 : len(spaces)/2]
			if terminalWidth > len(lineFormArrayOfOutput)+len(spacesLeft)+len(spacesRight) {
				spacesRight = spaces[0 : len(spaces)/2+1]
			} else if terminalWidth < len(lineFormArrayOfOutput)+len(spacesLeft)+len(spacesRight) {
				spacesRight = spaces[0 : len(spaces)/2-1]
			}
			fmt.Print(spacesLeft, lineFormArrayOfOutput, spacesRight)
		}
	case "justify":
		counterForSpaceInText := SpaceCount(text)
		for _, lineFormArrayOfOutput := range arrayOfOutput {
			spaces := CreateSpaces(lineFormArrayOfOutput, terminalWidth)
			spacesHalf := spaces[0 : len(spaces)/(counterForSpaceInText)]
			for i := 0; i < len(lineFormArrayOfOutput); i++ {
				if lineFormArrayOfOutput[i] == '+' {
					fmt.Print(spacesHalf)
				} else {
					fmt.Print(string(lineFormArrayOfOutput[i]))
				}
			}
			fmt.Println()
		}
	}
}


func Justify() {
	align := internal.FindFile(os.Args[1])
	if align == "Failed" {
		fmt.Println("Missing argument")
		os.Exit(0)
	}

	var text string
	var banner string

	if len(os.Args) == 4 {
		text = os.Args[2]
		banner = os.Args[3]
	} else {
		text = os.Args[2]
		banner = "standard"
	}

	text, notValidText := internal.CheckIsAscii(text)
	if notValidText != "" {
		fmt.Println(notValidText)
		os.Exit(0)
	}

	if !internal.CheckIsBanner(banner) {
		fmt.Println("Wrong number of arguments\nUsage: \"go run . --align=right something shadow\"")
		os.Exit(0)
	}

	if !internal.CheckForChangeFile("assets/"+banner+".txt", banner) {
		fmt.Println("The file has been changed , the program will close")
		os.Exit(0)
	}

	text = strings.Trim(text, " ")
	splitedWords, standardAscii := internal.PrepareForOutput(banner, text)
	output := ""

	for _, word := range splitedWords {
		if word == "" {
			output += "\n"
			continue
		}
		for index := 1; index <= 8; index++ {
			check := true
			for _, ch := range word {
				if ch != ' ' {
					check = true
				}
				if ch == ' ' && align == "justify" && check {
					output += "+"
					check = false
				}
				if !check {
					continue
				}
				output += standardAscii[int((ch-32)*9)+index]
			}
			output += "\n"
		}
	}
	Aling(align, output, text)
}