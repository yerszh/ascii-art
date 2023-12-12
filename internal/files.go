package asciiArt

import (
	"fmt"
	"os"
)

func FindFile(firstArgument string) string {
	textAsFileName := ""
	for i := len(firstArgument) - 1; i > 0; i-- {
		if firstArgument[i] != '=' {
			textAsFileName = string(firstArgument[i]) + textAsFileName
		} else {
			break
		}
	}
	return textAsFileName
}

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	return file, err
}

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	return file, err
}

// Функция для записи в файла
func WtiteFile(fileName *os.File, graphicRepresentation string) (string, error) {
	fileName.Seek(0, 0)      // Перемещение указателя в начало файла			\
	fileName.Truncate(0)     // Очистка файла путем урезания до нулевой длины	 > Очистка файла от предыдущей записи
	fileName.WriteString("") // Запись пустой строки							/

	_, err := fileName.WriteString(graphicRepresentation) // Обработка ошибки при записи данных в файл
	if err != nil {                                       // Если ошибка при записи данных в файл
		return "Unable to write to file", err // Возврат результата функции с ошибкой
	}

	defer fileName.Close() // Закрытие файла

	return "", err // Возврат результата функции с положительным результатом
}

// Функция для подготовки файла
func PrepareFile() *os.File {
	textAsFileName := FindFile(os.Args[1]) // Переменная для результата функции
	file, err := OpenFile(textAsFileName)  //	Переменная для хранения имени файла после открытия
	if err != nil {                        //	Если ошибка
		fmt.Println("Could not open the file ") // Вывод сообщение об ошибке при открытии

		file, err = CreateFile(textAsFileName) // Переменная для хранения имени файла после создания

		if err != nil { // Если ошибка
			fmt.Println("Could not create the file ") // Вывод сообщение об ошибке при создании

			os.Exit(1) // Выход из программы с ошибкой 1
		} else { // Если ошибки нет
			fmt.Println("File created") //	Вывод сообщение об успешном создании
		}
	}

	return file
}
