package asciiArt

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	internal "ascii-art-output/internal"
)

// Функция для подсчета пробелов в тексте
func SpaceCount(text string) int {
	counterForSpaceInText := 0 // Создание переменной с 0 значением

	for i := 0; i < len(text); i++ { // Цикл для перебора символов в тексте
		if i != 0 && i <= len(text)-2 {
			if text[i-1] != ' ' && text[i] == ' ' { // Если символ равен " "
				counterForSpaceInText++ // Увеличение переменной на 1
			}
		}
	}

	return counterForSpaceInText // Возвращение значения
}

// Функция для создания пробелов в ширину терминала
func CreateSpaces(lineFormArrayOfOutput string, terminalWidth int) string {
	spaces := "" // Создание пустой переменной для будущих пробелов

	for terminalWidth-len(lineFormArrayOfOutput) > 0 { // Цикл пока в терминале есть место для пробелов
		spaces += " "   // Создание пробела
		terminalWidth-- // Ширина терминала -1
	}

	return spaces // Возвращение значения с пробелами
}

// Функция для определения ширины терминала
func GetTermWidth() int {
	out, er1 := exec.Command("tput", "cols").Output() // Выполнить команду "tput cols" и получить результат
	out1 := strings.TrimSpace(string(out))            // Удалить лишние пробелы из вывода
	num, er2 := strconv.Atoi(out1)                    // Преобразовать строку в целое число

	if er1 != nil { // Проверить ошибки при выполнении команды "tput cols"
		os.Exit(0)
	}

	if er2 != nil { // Проверить ошибки при преобразовании строки в число
		os.Exit(0)
	}

	return num // Возвращение значение ширины терминала
}

// Функция для применения флага
func Aling(align string, output string, text string) {
	arrayOfOutputWrong := strings.Split(output, "\n")                  // Разделение входящей строки и преобразование ее в массив
	arrayOfOutput := arrayOfOutputWrong[0 : len(arrayOfOutputWrong)-1] // Удаление последнего элемента
	terminalWidth := GetTermWidth()                                    // СОхранение результата функции
	countSpace := CreateSpaces(arrayOfOutput[0], terminalWidth)        // Добавление пробелов в строки

	if align == "justify" && SpaceCount(text) == 0 {
		align = "center"
	}

	if terminalWidth < len(arrayOfOutput[0])+len(countSpace) {
		fmt.Println("Terminal width is too small") // Вывод сообщение об ошибке при создании

		os.Exit(0) // Выход из программы с ошибкой 1
	}

	switch align { // Выбор между флагами
	case "left": // Если равен left
		fmt.Print(output)
	case "right":
		for _, lineFormArrayOfOutput := range arrayOfOutput { // Цикл для перебора строк в массиве
			// fmt.Println(len(lineFormArrayOfOutput))
			spacesLeft := CreateSpaces(lineFormArrayOfOutput, terminalWidth) // Добавление пробелов в строки

			fmt.Print(spacesLeft, lineFormArrayOfOutput) // Вывод строки
		}
	case "center":
		for _, lineFormArrayOfOutput := range arrayOfOutput { // Цикл для перебора строк в массиве
			spaces := CreateSpaces(lineFormArrayOfOutput, terminalWidth) // Добавление пробелов в строки
			spacesLeft := spaces[0 : len(spaces)/2]                      // Разделение пробелов пополам
			spacesRight := spaces[0 : len(spaces)/2]                     // Создание пустой переменной для пробелов справка от графического предствления

			if terminalWidth > len(lineFormArrayOfOutput)+len(spacesLeft)+len(spacesRight) { // Если длина строки нечетная
				spacesRight = spaces[0 : len(spaces)/2+1] // Добавление пробела справа
			} else if terminalWidth < len(lineFormArrayOfOutput)+len(spacesLeft)+len(spacesRight) { // Если пробелов четное кол-во
				spacesRight = spaces[0 : len(spaces)/2-1] // Добавление пробела справа
			}
			// fmt.Println(terminalWidth, len(lineFormArrayOfOutput), len(spacesLeft), len(spacesRight))
			fmt.Print(spacesLeft, lineFormArrayOfOutput, spacesRight) // Вывод строки
		}
	case "justify":
		counterForSpaceInText := SpaceCount(text)
		for _, lineFormArrayOfOutput := range arrayOfOutput { // Цикл для перебора строк в массиве
			spaces := CreateSpaces(lineFormArrayOfOutput, terminalWidth) // Добавление пробелов в строки
			spacesHalf := spaces[0 : len(spaces)/(counterForSpaceInText)]
			// fmt.Println(terminalWidth, len(spacesHalf), len(lineFormArrayOfOutput))
			for i := 0; i < len(lineFormArrayOfOutput); i++ {
				if lineFormArrayOfOutput[i] == '+' {
					fmt.Print(spacesHalf)
				} else {
					fmt.Print(string(lineFormArrayOfOutput[i]))
				}
			}

			fmt.Println()
		}
	}
}

func Justify() {
	align := internal.FindFile(os.Args[1]) // Сохранение значения 1 аргумента как флаг

	if align == "Failed" { // Если не прописан флаг и его аргумент
		fmt.Println("Missing argument") // Вывод сообщение о пропущенном аргументе

		os.Exit(0) // Выход из программы с ошибкой 1
	}

	var text string   // Переменная для хранения текста
	var banner string // Переменная для хранения баннера

	if len(os.Args) == 4 { //	Если кол-во аргументов 4
		text = os.Args[2]   //  Сохранение значения 3 аргумента как текст
		banner = os.Args[3] //	Сохранение значения 4 аргумента как баннер
	} else { //	Если кол-во аргументов 3
		text = os.Args[2]
		banner = "standard" //	Изменение значения на standard
	}

	text, notValidText := internal.CheckIsAscii(text) // Переменная для хранения текста

	if notValidText != "" { // Если нет валидного текста
		fmt.Println(notValidText) // Вывод сообщения об ошибки

		os.Exit(0) // Выход из программы
	}

	if !internal.CheckIsBanner(banner) {
		fmt.Println("Wrong number of arguments\nUsage: \"go run . --align=right something shadow\"") // Вывод сообщения

		os.Exit(0)
	}

	if !internal.CheckForChangeFile("assets/"+banner+".txt", banner) { // Если файл изменен
		fmt.Println("The file has been changed , the program will close") // Вывод сообщения об изменении в файле

		os.Exit(0) // Выход из программы с ошибкой 1
	}

	text = strings.Trim(text, " ")

	splitedWords, standardAscii := internal.PrepareForOutput(banner, text) // Сохранение результата функции
	output := ""                                                           // Создание пустой переменной для сохранения графического рисунка

	for _, word := range splitedWords { // Проходим по разбитым словам и отображаем их в виде ASCII-арт.
		if word == "" { //  Если равно пустоте
			output += "\n" // Добавление новой строки

			continue //	Следующий элемент
		}
		for index := 1; index <= 8; index++ { // Цикл для отображения 8 строк текста ввиде графического ключа
			check := true

			for _, ch := range word { //	Цикл, для каждой буквы текста
				if ch != ' ' {
					check = true
				}
				if ch == ' ' && align == "justify" && check {
					output += "+"
					check = false
				}

				if !check {
					continue
				}
				output += standardAscii[int((ch-32)*9)+index] // Сохранение графического ключа построчно
			}
			output += "\n"
		}
	}

	Aling(align, output, text) // Вызов функции
}
