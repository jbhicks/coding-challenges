package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// PrettyPrint takes an interface{} and prints it in a pretty format.
func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

// Gets the string content of a file defined by the path parameter
func GetFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Reusable error checking function
func Check(e error, atTheDisco bool) {
	if e != nil {
		log.Fatal(e)
		if atTheDisco {
			panic(e)
		}
	}
}
