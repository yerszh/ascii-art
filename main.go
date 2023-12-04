package main

import (
	"fmt"
	"os"
)

type AsciiArt struct {
	runeArrayOfText []rune 
	text            string 
	banner          string 
}


func (ascii *AsciiArt) CheckForArgs(args []string) string {
	if len(args) > 3 { 
		return "Wrong number of arguments" 
	} else if len(args) == 3 { 
		switch os.Args[1] { 
		case "standard": 
			ascii.banner = "standard" 
		case "shadow": 
			ascii.banner = "shadow" 
		case "thinkertoy": 
			ascii.banner = "thinkertoy"
			default: 
			switch args[2] { 
			case "standard":
				ascii.banner = "standard"
			case "shadow":
				ascii.banner = "shadow"
			case "thinkertoy":
				ascii.banner = "thinkertoy"
			default:
				return "Wrong arguments"
			}

			ascii.text = args[1]
		}
	} else if len(args) == 2 {
		ascii.banner = "standard"
		ascii.text = args[1] 
	} else if len(args) == 1 {
		return "Wrong arguments"
	}

	if ascii.text == "" { 
		ascii.text = args[2] 
	}

	textAfterCheckForASCII := "" 

	for i := 0; i < len(ascii.text); i++ { 
		if ascii.text[i] >= 0 && ascii.text[i] <= 126 { 
			textAfterCheckForASCII = textAfterCheckForASCII + string(ascii.text[i]) 
		}
	}

	ascii.text = textAfterCheckForASCII 

	return "Pass" 
}

func main() {
	var ascii AsciiArt 

	if ascii.CheckForArgs(os.Args) != "Pass" { 
		return 
	}

	ascii.banner = "assets/" + ascii.banner + ".txt" 

	for _, char := range ascii.text { 
		ascii.runeArrayOfText = append(ascii.runeArrayOfText, rune(char)) 
	}

	_, err := ascii.OpenFile() 

	if err != nil { 
		fmt.Println("The file could not be opened, the program will be closed", err) 
	} else {
		ascii.ReadFile() 
	}
}