package util

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint takes an interface{} and prints it in a pretty format.
func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
