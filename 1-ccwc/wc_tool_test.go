package main

import (
	"os"
	"testing"
)

func TestByteSize(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("failed to open test.txt: %v", err)
	}
	defer file.Close()

	var expected int64 = 342190
	result := getFileByteSize(file)
	if result != expected {
		t.Errorf("getFileByteSize() = %d; want %d", result, expected)
	}
}

func TestLineCount(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("failed to open test.txt: %v", err)
	}
	defer file.Close()

	var expected int = 7145
	result := lineCount(file)
	if result != expected {
		t.Errorf("lineCount() = %d; want %d", result, expected)
	}
}

func TestWordCount(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("failed to open test.txt: %v", err)
	}
	defer file.Close()

	var expected int = 58164
	result := wordCount(file)
	if result != expected {
		t.Errorf("wordCount() = %d; want %d", result, expected)
	}
}

func TestCharacterCount(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("failed to open test.txt: %v", err)
	}
	defer file.Close()

	expected := 342190
	result := characterCount(file)
	if result != expected {
		t.Errorf("characterCount() = %d; want %d", result, expected)
	}
}
