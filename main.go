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

	
	
	default: 
		option.Fs()
	}
} else {
	fmt.Println("Wrong number of arguments") 
}
}
