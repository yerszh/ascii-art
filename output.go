package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)


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
	file, err := os.OpenFile(ascii.banner, os.O_RDONLY, 0666) 
	return file, err 
}


func onlyNewLines(input []string) bool {
	for _, ch := range input { 
		if ch != "" { 
			return false 
		}
	}
	return true 
}

func (ascii *AsciiArt) ReadFile() {
	dataBytes, _ := os.ReadFile(ascii.banner) 
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

		if ascii.banner == "assets/thinkertoy.txt" { 

			standardAscii = strings.Split(string(dataBytes), "\r\n") 
		} else {
			standardAscii = strings.Split(string(dataBytes), "\n") 
		}

		inputWords := ascii.text                                            
		splittedWords := regexp.MustCompile(`\n|\\n`).Split(inputWords, -1) 

		if onlyNewLines(splittedWords) { 

			splittedWords = splittedWords[1:] 
		}

		for _, word := range splittedWords { 
			if word == "" {
				fmt.Println()
				continue
			}
			for index := 1; index <= 8; index++ {
				for _, ch := range word {
					fmt.Print(standardAscii[int((ch-32)*9)+index]) 
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("The file has been changed , the program will close")
	}
}
