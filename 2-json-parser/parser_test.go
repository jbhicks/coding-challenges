package main

import (
	"bytes"
	"log"
	"os/exec"
	"testing"
)

func TestMainFunctionPart1(t *testing.T) {
	t.Run("Invalid JSON", func(t *testing.T) {
		cmd := exec.Command("go", "run", ".", "tests/step1/invalid.json")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		err := cmd.Run()
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 1 {
				// Expected exit status 1, test passes
				log.Printf("Expected failure: %s", out.String())
				return
			}
		}

		if err != nil {
			t.Fatalf("Unexpected error: %v\nOutput: %s", err, out.String())
		}

		t.Fatalf("Expected exit status 1 but got success\nOutput: %s", out.String())
	})

	t.Run("Valid JSON", func(t *testing.T) {
		cmd := exec.Command("go", "run", ".", "tests/step1/valid.json")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		err := cmd.Run()
		if err != nil {
			t.Fatalf("Expected success but got error: %v\nOutput: %s", err, out.String())
		}

		log.Printf("Output: %s", out.String())
	})
}
