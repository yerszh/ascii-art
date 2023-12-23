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
	case "--output": //	Если флаг равен output
		if len(os.Args) == 4 || len(os.Args) == 3 { // Если кол-во аргументов 4 или 3
			option.Output() //  Вызов функции
		} else { // Если кол-во аргументов не 4 и не 3
			fmt.Println("Wrong number of arguments\nUsage: \"go run . --output=<fileName.txt> something standard\"") // Вывод сообщения
		}
		
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
