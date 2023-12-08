package asciiArt

import (
	internal "ascii-art-output/internal"
	"fmt"
	"os"
)

func Fs() {
	var banner string  
	text := os.Args[1] 

	if len(os.Args) == 3 { 
		banner = os.Args[2] 
	} else if len(os.Args) == 2 { 
		banner = "standard" 
	} else {
		fmt.Println("Wrong number of arguments") 
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
					fmt.Print(standardAscii[int((ch-32)*9)+index])
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("The file has been changed , the program will close") 
		os.Exit(1)
	}
}
