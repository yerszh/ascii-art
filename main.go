package main

import (
	internal "ascii-art-output/internal"
	option "ascii-art-output/option"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 { 
		if len(os.Args) == 2 {
			option.Fs()

			return
		}
		
		if os.Args[2] == "" {
			return
		}

		flag := internal.ClearFlag(os.Args[1]) 

		switch flag {
		case "--output": 
			if len(os.Args) == 4 || len(os.Args) == 3 { 
				option.Output() 
			} else { 
				fmt.Println("Wrong number of arguments\nUsage: \"go run . --output=<fileName.txt> something standard\"") 
			}
		case "--reverse": 
		if len(os.Args) == 2 {
			option.Reverse()
		} else { 
			fmt.Println("Wrong number of arguments\nUsage: \"go run . --reverse=<fileName.txt>\"") 
		}	
		case "--color": 
			if len(os.Args) >= 3 && len(os.Args) <= 7 {
				option.Color()
			} else {
				fmt.Println("Wrong number of arguments\nUsage: \"go run . --color=rgb(48, 255, 229) smthng something thinkertoy\"") 
			}
		case "--align": 
			if len(os.Args) == 4 || len(os.Args) == 3 { 
				option.Justify()
			} else {
				fmt.Println("Wrong number of arguments\nUsage: \"go run . --align=right something shadow\"") 
			}
		}
	} else {
		fmt.Println("Wrong number of arguments") 
	}
}
