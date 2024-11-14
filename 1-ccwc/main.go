package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, playground", os.Args)
	// read the number of bytes contained in file.txt
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileInfo.Size())
	}

}
