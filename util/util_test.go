package util

import (
	"os"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	// Write some content to the temporary file
	content := "Hello, World!"
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Test GetFileContent function
	result, err := GetFileContent(tmpfile.Name())
	Check(err, false)
	if result != content {
		t.Errorf("Expected %q, got %q", content, result)
	}

	// Test GetFileContent with a non-existent file
	result, _ = GetFileContent("non_existent_file.txt")
	if result != "" {
		t.Errorf("Expected empty string for non-existent file, got %q", result)
	}
}
