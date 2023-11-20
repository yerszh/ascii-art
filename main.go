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
	if len(args) > 3 { //	Если Аргументов больше 3 или 1
		return "Wrong number of arguments" //	Возврщения сообщения об ошибке
	} else if len(args) == 3 { //	Если аргументов 3
		switch os.Args[1] { //	Выбор для первого аргумента
		case "standard": //	Если первый аргумент равен "standard"
			ascii.banner = "standard" //	Изменения значения переменной на "standard"
		case "shadow": //	Если первый аргумент равен "sahdow"
			ascii.banner = "shadow" //	//	Изменения значения переменной на "sahdow"
		case "thinkertoy": //	Если первый аргумент равен "thinkertoy"
			ascii.banner = "thinkertoy" //	//	Изменения значения переменной на "thinkertoy"
		default: //	Если нет совпадения
			

			ascii.text = args[1] //	Изменения значения переменной на совпадения в операторе выбора
		}
	} else if len(args) == 2 {
		ascii.banner = "standard"
		ascii.text = args[1] // Изменение значения на аргумент
	} else if len(args) == 1 {
		return "Wrong arguments"
	}

	if ascii.text == "" { //	Если переменная текста пустая
		ascii.text = args[2] //	Изменения значения переменной на первый аргумент
	}

	textAfterCheckForASCII := "" // Создание переменной для проверки символом по таблице ASCII

	for i := 0; i < len(ascii.text); i++ { // Цикл для проверки символом по таблице ASCII
		if ascii.text[i] >= 0 && ascii.text[i] <= 126 { // Если символ имеет номер от 0 до 126
			textAfterCheckForASCII = textAfterCheckForASCII + string(ascii.text[i]) // Сохранение символа
		}
	}

	ascii.text = textAfterCheckForASCII // Сохранение символов,которые прошли проверку

	return "Pass" //	Возвращение сообщение об успешной проверке
}

func main() {
	var ascii AsciiArt //	Создания переменной структуры

	if ascii.CheckForArgs(os.Args) != "Pass" { //	Если возвращенное значение из функции неравно "Pass"
		return //	Выход из программы
	}

	ascii.banner = "assets/" + ascii.banner + ".txt" //	Конкатинация строк

	for _, char := range ascii.text { //	Цикл для перебора текста
		ascii.runeArrayOfText = append(ascii.runeArrayOfText, rune(char)) //	Добавление в срез текст в представлении ввиде рун
	}

	_, err := ascii.OpenFile() // Создание переменной для файла cо шрифтом

	if err != nil { // Если есть ошибка при открытии файла с текстом
		//  > Обработка ошибки при открытии файла с текстом и вывод сообщения
		fmt.Println("The file could not be opened, the program will be closed", err) //||\\
		fmt.Println("Создание файла - ", ascii.banner) // Вывод сообщения о попытке создани файла

		
		if err != nil { // Если есть ошибка при создании файла cо шрифтом
			fmt.Println("Ошибка при создании файла", err) // Обработка ошибки при создании файла cо шрифтом

			return // Выход из программы
		} else {
			fmt.Println("Файл создан") // Вывод сообщения об успешном cоздании файла

		}
	} else {
		ascii.ReadFile() // Вызов функции
	}
}
