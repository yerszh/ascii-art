package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Функция для создания файла
/*func (ascii *AsciiArt) CreateFile() (*os.File, error) {
	file, err := os.Create(ascii.banner) // Создание файла

	return file, err // Возвращение файла и ошибки для дальнейшей обработки
}*/

// Функция для открытия файла
func FileMD5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (ascii *AsciiArt) OpenFile() (*os.File, error) {
	file, err := os.OpenFile(ascii.banner, os.O_RDONLY, 0666) // Открытие файла и для чтения и для записи

	return file, err // Возвращение файла и ошибки для дальнейшей обработки
}

// Функция для определения новой строки
func onlyNewLines(input []string) bool {
	for _, ch := range input { // Цикл для перебора аргумента для вывода
		if ch != "" { // Проверка на пустую строку
			return false // Возвращаем false
		}
	}
	return true // Возвращаем true
}

// Функция для чтения из файла и вывода на экран
func (ascii *AsciiArt) ReadFile() {
	dataBytes, _ := os.ReadFile(ascii.banner) // Считываем содержимое файла в байтовый массив.
	checkHashSumFile := FileMD5(ascii.banner)
	filePass := true

	switch ascii.banner {
	case "assets/standard.txt":
		if "ac85e83127e49ec42487f272d9b9db8b" != checkHashSumFile {
			filePass = false
		}
	case "assets/thinkertoy.txt":
		if "86d9947457f6a41a18cb98427e314ff8" != checkHashSumFile {
			filePass = false
		}
	case "assets/shadow.txt":
		if "a49d5fcb0d5c59b2e77674aa3ab8bbb1" != checkHashSumFile {
			filePass = false
		}
	}

	if filePass {
		standardAscii := []string{}

		if ascii.banner == "assets/thinkertoy.txt" { // Проверяем, является ли файл "assets/thinkertoy.txt".

			standardAscii = strings.Split(string(dataBytes), "\r\n") // Разделяем текст по символам перевода строки Windows.
		} else {
			standardAscii = strings.Split(string(dataBytes), "\n") // Разделяем текст по обычным символам перевода строки.
		}

		inputWords := ascii.text                                            // Получаем входные слова для преобразования в ASCII-арт.
		splittedWords := regexp.MustCompile(`\n|\\n`).Split(inputWords, -1) // Разбиваем входные слова на части, используя регулярное выражение.

		if onlyNewLines(splittedWords) { // Проверяем, состоит ли текст только из символов перевода строки.

			splittedWords = splittedWords[1:] // Если да, удаляем первый элемент (пустую строку).
		}

		for _, word := range splittedWords { // Проходим по разбитым словам и отображаем их в виде ASCII-арт.
			if word == "" {
				fmt.Println()
				continue
			}
			for index := 1; index <= 8; index++ {
				for _, ch := range word {
					fmt.Print(standardAscii[int((ch-32)*9)+index]) // Отображаем соответствующий символ ASCII из стандартного набора.
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("The file has been changed , the program will close")
	}
}
