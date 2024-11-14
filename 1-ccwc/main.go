package main

import (
	"bufio"
	"fmt"
	"os"

	util "github.com/jbhicks/coding-challenges-go/util"
)

func main() {
	// open file defined by Args[2]
	f, _ := os.Open(os.Args[2])
	defer f.Close()

	util.PrettyPrint(f)

	switch os.Args[1] {
	case "-c":
		fmt.Println(getFileByteSize(f), os.Args[2])
	case "-l":
		fmt.Println(lineCount(f), os.Args[2])
	case "-w":
		fmt.Println(wordCount(f), os.Args[2])
	case "-m":
		fmt.Println("-m switch")
	default:
		fmt.Println("default switches")
	}
}

func wordCount(f *os.File) int {
	// scan the file for words and count them
	count := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}

func lineCount(f *os.File) int {
	// scan the file for words and count them
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count++
	}

	return count
}

func getFileByteSize(f *os.File) int64 {
	fileInfo, err := f.Stat()
	if err != nil {
		return -1
	}
	return fileInfo.Size()
}
