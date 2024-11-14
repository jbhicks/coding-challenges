package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "-c":
		fmt.Println(getFileByteSize(os.Args[2]), os.Args[2])
	case "-w":
		fmt.Println(wordCount(os.Args[2]), os.Args[2])
	case "-m":
		fmt.Println("-m switch")
	default:
		fmt.Println("default switches")
	}
}

func wordCount(s string) int {
	// open file specified by Args[2]
	f, err := os.Open(s)
	if err != nil {
		return -1
	}
	defer f.Close()

	// scan the file for words and count them
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count++
	}

	return count
}

func getFileByteSize(fileName string) int64 {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return -1
	}
	return fileInfo.Size()
}
