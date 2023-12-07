package asciiArt

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

//	Функция для удаления лишнего и сохранение флага из аргумента
func ClearFlag(firstArgument string) string { //
	flag := "" //	Создание пустой переменной для хранения флага								 					  \

	for _, char := range firstArgument { //	Перебор по символам первого аргумента										\
		if char != '=' { //	Если символ неравен =
			flag += string(char) //	Конкатинация посимвольно с первого аргумента в переменную для хранения флага		  >  УДАЛЕНИЕ ЛИШЕНГО И СОХРАНЕНИЕ ФЛАГА ИЗ ПЕРВОГО АРГУМЕНТА
		} else { //	Если символ равен =
			break //	Выход из цикла																					/
		}
	} //																											  /

	return flag //	Возвращение переменной с флагом
}

// Функция для хеширования содержимого файла
func FileMD5(path string) string {
	h := md5.New() // Создаем новый экземпляр хеш-функции MD5

	f, err := os.Open(path) // Открываем файл по заданному пути
	if err != nil {         // Если возникла ошибка при открытии файла
		os.Exit(1) // Выход из программы с ошибкой 1
	}

	defer f.Close() // Закрываем файл после завершения функции

	_, err = io.Copy(h, f) // Копируем содержимое файла в хеш-функцию MD5
	if err != nil {        // Если возникла ошибка при копировании
		os.Exit(1) // Выход из программы с ошибкой 1
	}

	return fmt.Sprintf("%x", h.Sum(nil)) // Возвращаем значение MD5-хеша в виде строки в шестнадцатеричном формате
}

// Функция для проверки на изменения в файлах со шрифтами
func CheckForChangeFile(textAsFileName string, banner string) bool {
	checkHashSumFile := FileMD5(textAsFileName) //	Сохранение результата функции

	switch banner { //	Выбор для 3 аргумента
	case "standard": //	Если равен standard
		if "ac85e83127e49ec42487f272d9b9db8b" != checkHashSumFile { // Если результат функции равен готовому хешу
			return false // Возвращение false
		}
	case "thinkertoy":
		if "86d9947457f6a41a18cb98427e314ff8" != checkHashSumFile {
			return false
		}
	case "shadow":
		if "a49d5fcb0d5c59b2e77674aa3ab8bbb1" != checkHashSumFile {
			return false
		}
	}

	return true // Возвращение trues
}
