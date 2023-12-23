package asciiArt

import (
	"fmt"
	"os"

	internal "ascii-art-output/internal"
)


func Output() {
	file := internal.PrepareFile() 

	var banner string 

	

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
