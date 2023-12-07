package main

import (
	internal "ascii-art-output/internal"
	option "ascii-art-output/option"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 { // Если кол-во аргументов больше или равно 2
		flag := internal.ClearFlag(os.Args[1]) // Сохранение результата функции

		switch flag { // Выбор между результатом функции для определения опционалки
		
		default: // Если совпадения не было
			option.Fs()
		}
	} else {
		fmt.Println("Wrong number of arguments") // Вывод сообщения
	}
}
