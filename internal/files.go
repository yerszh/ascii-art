package asciiArt

import (
	"fmt"
	"io"
	"os"
)


func ReadFile(fileName *os.File) string {
dataByteFromFile := make([]byte, 64)         
textFromFile, textAfterCheckForASCII := "", "" 

for { 
	n, err := fileName.Read(dataByteFromFile)
	if err == io.EOF {                   
		break 
	}
	textFromFile = textFromFile + string(dataByteFromFile[:n]) 
}

for i := 0; i < len(textFromFile); i++ { 
	if textFromFile[i] >= 0 && textFromFile[i] <= 126 { 
		textAfterCheckForASCII = textAfterCheckForASCII + string(textFromFile[i]) 
	}
}

defer fileName.Close() 

return textAfterCheckForASCII 
}




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

func WtiteFile(fileName *os.File, graphicRepresentation string) (string, error) {
	fileName.Seek(0, 0)
	fileName.Truncate(0)
	fileName.WriteString("")
	_, err := fileName.WriteString(graphicRepresentation)
	if err != nil {
		return "Unable to write to file", err
	}
	defer fileName.Close()
	return "", err
}

func PrepareFile() *os.File {
	textAsFileName := FindFile(os.Args[1])
	file, err := OpenFile(textAsFileName)
	if err != nil {
		fmt.Println("Could not open the file")
		file, err = CreateFile(textAsFileName)
		if err != nil {
			fmt.Println("Could not create the file")
			os.Exit(1)
		} else {
			fmt.Println("File created")
		}
	}
	return file
}