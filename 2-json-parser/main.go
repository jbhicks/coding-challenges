package main

import (
	"log"
	"os"

	"github.com/jbhicks/coding-challenges-go/util"
)

func main() {
	fileName := os.Args[1]
	log.Println("fileName:", fileName)

	// open the file in the first argument
	fileContent := util.GetFileContent(os.Args[1])

	if len(fileContent) == 0 {
		log.Fatal("File is empty")
	}
}
