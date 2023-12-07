package asciiArt

import (
	"os"
	"regexp"
	"strings"
)

func onlyNewLines(input []string) bool {
	for _, ch := range input { // Цикл для перебора аргумента для вывода
		if ch != "" { // Проверка на пустую строку
			return false // Возвращаем false
		}
	}
	return true // Возвращаем true
}

func PrepareForOutput(banner string, text string) ([]string, []string) {
	dataBytes, _ := os.ReadFile("assets/" + banner + ".txt") // Считываем содержимое файла в байтовый массив.
	standardAscii := []string{}

	if banner == "thinkertoy" { // Проверяем, является ли файл "assets/thinkertoy.txt".
		standardAscii = strings.Split(string(dataBytes), "\r\n") // Разделяем текст по символам перевода строки Windows.
	} else {
		standardAscii = strings.Split(string(dataBytes), "\n") // Разделяем текст по обычным символам перевода строки.
	}

	inputWords := text                                                  // Получаем входные слова для преобразования в ASCII-арт.
	splittedWords := regexp.MustCompile(`\n|\\n`).Split(inputWords, -1) // Разбиваем входные слова на части, используя регулярное выражение.

	if onlyNewLines(splittedWords) { // Проверяем, состоит ли текст только из символов перевода строки.
		splittedWords = splittedWords[1:] // Если да, удаляем первый элемент (пустую строку).
	}

	return splittedWords, standardAscii
}
