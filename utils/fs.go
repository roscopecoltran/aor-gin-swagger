package utils

import (
	"fmt"
	"os"
)

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("error: ", err)
		return false, err
	}
	fmt.Printf("IsDir? %t , dir=%s\n", fileInfo.IsDir(), path)
	fmt.Println("fileInfo: ", fileInfo)
	return fileInfo.IsDir(), err
}
