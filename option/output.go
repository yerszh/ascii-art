package asciiArt

import (
	"fmt"
	"os"

	internal "ascii-art-output/internal"
)


func Output() {
	file := internal.PrepareFile() 

	var banner string 

	if len(os.Args) == 4 { 
		banner = os.Args[3] 
	} else { 
		banner = "standard" 
	}

	if !internal.CheckIsBanner(banner) {
		fmt.Println("Wrong number of arguments\nUsage: \"go run . --color=48, 255, 229 smthng something thinkertoy\"") 

		os.Exit(1)
	}

	if internal.CheckForChangeFile("assets/"+banner+".txt", banner) { 
		splitedWords, standardAscii := internal.PrepareForOutput(banner, os.Args[2]) 
		output := ""                                                                

		for _, word := range splitedWords { 
			if word == "" { 
				output += "\n" 

				continue 
			}
			for index := 1; index <= 8; index++ {
				for _, ch := range word { 
					output += standardAscii[int((ch-32)*9)+index] 
				}
				output += "\n"
			}
		}

		internal.WtiteFile(file, output) 

		fmt.Println("Done") 
	} else {
		fmt.Println("The file has been changed , the program will close") 

		os.Exit(1)
	}
}
