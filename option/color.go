package asciiArt

import (
	"fmt"
	"os"
	"strings"

	internal "ascii-art-output/internal"
)

func Atoi(text string) int {
	result := 0 // Создание переменной для дальнейшего сохранения результата

	for _, char := range text { // Цикл для перебора элементов строки
		if char >= '0' && char <= '9' { // Если элемент строки цифра
			number := int(char - '0')   // Преобразование этой byte в int
			result = result*10 + number //  Сохранение цифры и умножение предыдущей сохраненой на 10
		}
	}

	return result // Возвращение преобразованного числа
}

// Функция для нахождения кодов
func FindRGB(color string) (int, int, int) {
	arrayOfRGBSystem := []string{}
	r, g, b := 0, 0, 0

	if color[len(color)-1] != ')' {
		r = Atoi(os.Args[1])
		g = Atoi(os.Args[2])
		b = Atoi(os.Args[3])
	} else {
		arrayOfRGBSystem = strings.Split(color, ",")

		if len(arrayOfRGBSystem) == 3 {
			r = Atoi(arrayOfRGBSystem[0])
			g = Atoi(arrayOfRGBSystem[1])
			b = Atoi(arrayOfRGBSystem[2])
		}
	}

	return r, g, b
}

// Функция для определения цвета
func PickColor(color string) string {
	switch color { // Выбор между указанным цветом
	case "black": // Если равно black
		return "\u001b[30m" // Возвращение значения по коду ansi
	case "red":
		return "\u001b[31m"
	case "green":
		return "\u001b[32m"
	case "yellow":
		return "\u001b[33m"
	case "blue":
		return "\u001b[34m"
	case "magenta":
		return "\u001b[35m"
	case "cyan":
		return "\u001b[36m"
	case "white":
		return "\u001b[37m"
	}

	return "Failed" // Если совпадения нет, то возвращения ошибки
}

func rgbToAnsi(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b) // Экранирующие коды ANSI
}

// Функция для опционалки output
func Color() {
	color := strings.ToLower(internal.FindFile(os.Args[1])) // Сохранение первого аргумента как увет и преобразования его в нижний регистр

	if color == "failed" { // Если не прописан флаг и его аргумент
		fmt.Println("Missing argument\nTry like this  --> \" --color=rgb(48, 255, 229)\"") // Вывод сообщение о пропущенном аргументе

		os.Exit(1) // Выход из программы с ошибкой 1
	}

	var ansiCode string
	var r, g, b int
	var letterToBeColored string // Переменная для хранения букв, которые нужно перекрасить
	var text string              // Переменная для хранения текста
	var banner string            // Переменная для хранения баннера

	if color[len(color)-1] == ',' {
		r, g, b = FindRGB(color)      // Сохранение результата функции
		ansiCode = rgbToAnsi(r, g, b) // Сохранение результата функции

		switch len(os.Args) { // Выбор между кол-вом аргументов
		case 7:
			letterToBeColored = os.Args[4] // Сохранения значения 2 аргумента как буквы, которые нужно перекрасить
			text = os.Args[5]              // Сохранение значения 3 аргумента как текст
			banner = os.Args[6]            // Сохранение значения 4 аргумента как баннер
		case 6: // Есди кол-во аргументов равно 5
			if internal.CheckIsBanner(os.Args[len(os.Args)-1]) {
				text = os.Args[4]
				banner = os.Args[5] // Сохранение значения 4 аргумента как баннер
			} else {
				letterToBeColored = os.Args[4] // Сохранения значения 2 аргумента как буквы, которые нужно перекрасить
				text = os.Args[5]              // Сохранение значения 3 аргумента как текст
				banner = "standard"
			}
		case 5:
			text = os.Args[4]
			banner = "standard"
		}
	} else {
		if color[len(color)-1] == ')' {
			r, g, b = FindRGB(color)
			ansiCode = rgbToAnsi(r, g, b) // Сохранение результата функции
		} else {
			ansiCode = PickColor(color)

			if ansiCode == "Failed" {
				fmt.Println("Wrong color\nChoose one of these colors --> black, red, gree, yellow, blue, magenta, cyan, white") // Вывод сообщения

				os.Exit(1)
			}
		}

		switch len(os.Args) { // Выбор между кол-вом аргументов
		case 5: // Есди кол-во аргументов равно 5
			letterToBeColored = os.Args[2] // Сохранения значения 2 аргумента как буквы, которые нужно перекрасить
			text = os.Args[3]              // Сохранение значения 3 аргумента как текст
			banner = os.Args[4]            // Сохранение значения 4 аргумента как баннер
		case 4:
			if os.Args[len(os.Args)-1] == "standard" || os.Args[len(os.Args)-1] == "shadow" || os.Args[len(os.Args)-1] == "thinkertoy" { // Если последний аргумент равен одному из значений
				text = os.Args[2]
				banner = os.Args[3]
			} else {
				letterToBeColored = os.Args[2]
				text = os.Args[3]
				banner = "standard" // Изменение значения на standard
			}
		case 3:
			text = os.Args[2]
			banner = "standard"
		}
	}

	resetCode := "\x1b[0m"                            // Удалить цвет
	text, notValidText := internal.CheckIsAscii(text) // Переменная для хранения текста

	if notValidText != "" { // Если нет валидного текста
		fmt.Println(notValidText) // Вывод сообщения об ошибки

		os.Exit(1) // Выход из программы
	}

	if !internal.CheckIsBanner(banner) {
		fmt.Println("Wrong number of arguments\nUsage: \"go run . --color=rgb(48, 255, 229) smthng something thinkertoy\"") // Вывод сообщения

		os.Exit(1)
	}

	if internal.CheckForChangeFile("assets/"+banner+".txt", banner) { // Если файл не изменен
		splitedWords, standardAscii := internal.PrepareForOutput(banner, text) // Сохранение результата функции

		for _, word := range splitedWords {
			if word == "" {
				fmt.Println() // Новая строка

				continue // Следующий элемент
			}

			for index := 1; index <= 8; index++ { // Цикл для отображения 8 строк текста ввиде графического ключа
				for _, ch := range word { //	Цикл, для каждой буквы текста
					colored := false // Переменная для определения, должен ли символ быть раскрашен

					if len(letterToBeColored) != 0 { // Если буквы для окраски есть
						for _, letter := range letterToBeColored { // Перебор букв для окраски
							if ch == letter { //	Если буквы совпадают
								colored = true //	Изменение значения на true
								break          //	Выход из цикла
							}
						}
					} else { // Если буквы для окраски нет
						colored = true
					}

					if colored { // Значение true
						fmt.Print(ansiCode, standardAscii[int((ch-32)*9)+index]) //	Вывод окрашенных строк, сверху=вниз
					} else {
						fmt.Print(resetCode, standardAscii[int((ch-32)*9)+index]) // Вывод обычных строк, сверху-вниз
					}
				}

				fmt.Println()
			}
		}
	} else {
		fmt.Println("The file has been changed , the program will close") // Вывод сообщения об изменении в файле

		os.Exit(1) // Выход из программы с ошибкой 1
	}
}
