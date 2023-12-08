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

	if internal.CheckForChangeFile("assets/"+banner+".txt", banner) { // Если файл не изменен
		splitedWords, standardAscii := internal.PrepareForOutput(banner, text) // Сохранение результата функции

		for _, word := range splitedWords { // Проходим по разбитым словам и отображаем их в виде ASCII-арт.
			if word == "" { //  Если равно пустоте
				fmt.Println() // Добавление новой строки

				continue //	Следующий элемент
			}
			for index := 1; index <= 8; index++ { // Цикл для отображения 8 строк текста ввиде графического ключа
				for _, ch := range word { //	Цикл, для каждой буквы текста
					fmt.Print(standardAscii[int((ch-32)*9)+index]) // Отображаем соответствующий символ ASCII из стандартного набора.
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("The file has been changed , the program will close") // Вывод сообщения об изменении в файле

		os.Exit(1) // Выход из программы с ошибкой 1
	}
}
