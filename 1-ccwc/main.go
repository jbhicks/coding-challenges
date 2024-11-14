package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open file defined by Args[2]
	var fileName string
	if len(os.Args) < 3 {
		fileName = os.Args[1]
	} else {
		fileName = os.Args[2]
	}

	f, _ := os.Open(fileName)
	defer f.Close()

	switch os.Args[1] {
	case "-c":
		fmt.Println(getFileByteSize(f), fileName)
	case "-l":
		fmt.Println(lineCount(f), fileName)
	case "-w":
		fmt.Println(wordCount(f), fileName)
	case "-m":
		fmt.Println(characterCount(f), fileName)
	default:
		size := getFileByteSize(f)
		line := lineCount(f)
		word := wordCount(f)

		fmt.Println(size, line, word, fileName)
	}
}

func characterCount(f *os.File) int {
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count += len(scanner.Text())
	}
	return count
}

func wordCount(f *os.File) int {
	count := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}

func lineCount(f *os.File) int {
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
