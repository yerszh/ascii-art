package asciiArt

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func ClearFlag(firstArgument string) string { 
	flag := "" 							 					  

	for _, char := range firstArgument {									
		if char != '=' { 
			flag += string(char) 
		} else { 
			break 
		}
	} 																											  

	return flag 
}


func FileMD5(path string) string {
	h := md5.New() 

	f, err := os.Open(path) 
	if err != nil {     
		os.Exit(1) 
	}

	defer f.Close() 

	_, err = io.Copy(h, f) 
	if err != nil {    
		os.Exit(1) 
	}

	return fmt.Sprintf("%x", h.Sum(nil)) 
}


func CheckForChangeFile(textAsFileName string, banner string) bool {
	checkHashSumFile := FileMD5(textAsFileName) 

	switch banner { 
	case "standard": 
		if "ac85e83127e49ec42487f272d9b9db8b" != checkHashSumFile { 
			return false 
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

	return true 
}
