package main

import (
	internal "ascii-art-output/internal"
	option "ascii-art-output/option"
	"fmt"
	"os"
)

func main() {
if len(os.Args) >= 2 {
	flag := internal.ClearFlag(os.Args[1]) 

	switch flag { 

	case "--color": 
		if len(os.Args) >= 3 && len(os.Args) <= 7 {
			option.Color()
		} else {
			fmt.Println("Wrong number of arguments\nUsage: \"go run . --color=rgb(48, 255, 229) smthng something thinkertoy\"") // Вывод сообщения
		}
	
	default: 
		option.Fs()
	}
} else {
	fmt.Println("Wrong number of arguments") 
}
}
