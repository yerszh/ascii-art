package asciiArt

import (
	internal "ascii-art-output/internal"
	"fmt"
	"os"
	"strings"
)

// Функция для определения баннера
func BannerFinder(arrayOfGraphicRepresentation []string) string {
	for _, lineFormArrayOfOutput := range arrayOfGraphicRepresentation { // Перебор массива построчно
		if strings.ReplaceAll(lineFormArrayOfOutput, " ", "") != "" { // Если строка не пустая
			for i := 0; i < len(lineFormArrayOfOutput); i++ { //	Перебор строки посимвольно
				if i+1 <= len(lineFormArrayOfOutput)-1 { //	Если при добавлении 1 к позиции нет выхода за пределы
					if lineFormArrayOfOutput[i] == '_' && lineFormArrayOfOutput[i+1] != '|' { // Если символ равен _ а следующий не равен |
						return "standard" //	Возвращение результата
					}
				}

				if lineFormArrayOfOutput[i] == 'o' { // Если символ равен o
					return "thinkertoy"
				}
			}
		}
	}

	return "shadow"
}

func reverseAscii(font []string, text []string, pos int, count int, start int) {
	if pos != len(text[count]) {
		if start >= len(font) {
			return
		}

		l := len(font[start])
		if pos+l <= len(text[count]) {
			if count < 7 {
				if text[count][pos:l+pos] == font[start+count] {
					reverseAscii(font, text, pos, count+1, start)
				} else {
					reverseAscii(font, text, pos, 0, start+9)
				}
			} else {
				r := ((start - 1) / 9) + 32
				fmt.Printf("%c", r)
				reverseAscii(font, text, pos+l, 0, 1)
			}
		} else {
			reverseAscii(font, text, pos, 0, start+9)
		}
	}
}

func Reverse() {
	file := internal.PrepareFile()
	graphicRepresentation := internal.ReadFile(file)

	splittedText := strings.Split(graphicRepresentation, "\n")
	if len(splittedText) == 1 {
		splittedText = strings.Split(graphicRepresentation, "\r\n")
	}

	if !internal.CheckForChangeFile("assets/"+BannerFinder(splittedText)+".txt", BannerFinder(splittedText)) { // Если файл изменен
		fmt.Println("The file has been changed , the program will close") // Вывод сообщения об изменении в файле

		os.Exit(0) // Выход из программы с ошибкой 1
	}

	fontContent, errF := os.ReadFile("assets/" + BannerFinder(splittedText) + ".txt")
	if errF != nil {
		fmt.Println(errF)
		os.Exit(0)
	}

	fontData := strings.Split(string(fontContent), "\r\n")
	if len(fontData) == 1 {
		fontData = strings.Split(string(fontContent), "\n")
	}

	if len(splittedText) > 9 {
		for i := 0; i < len(splittedText)-1; {
			if len(splittedText[i]) > 0 {
				reverseAscii(fontData, splittedText[i:i+8], 0, 0, 1)
				fmt.Println()
				i = i + 8
			} else {
				fmt.Println()
				i = i + 1
			}
		}
	} else {
		reverseAscii(fontData, splittedText, 0, 0, 1)
		fmt.Println()
	}
}
