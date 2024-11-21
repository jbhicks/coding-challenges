package main

import (
	"log"
	"os"

	"github.com/jbhicks/coding-challenges-go/util"
)

func main() {
	fileName := os.Args[1]
	log.Println("Parse invoked with fileName:", fileName)

	// open the file in the first argument
	fileContent := util.GetFileContent(os.Args[1])

	if len(fileContent) == 0 {
		log.Fatal("File is empty")
	}

	// Verify first and last characters are '{' and '}' respectively
	if fileContent[0] != '{' || fileContent[len(fileContent)-1] != '}' {
		log.Fatal("Invalid JSON")
	} else {
		log.Println("Valid JSON")
		os.Exit(0)
	}
}
