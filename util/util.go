package util

import (
	"fmt"
	"reflect"
	"strings"
)

// PrettyPrint takes an interface{} and prints it in a pretty format.
func PrettyPrint(v interface{}) {
	value := reflect.ValueOf(v)
	fmt.Println(prettyPrintValue(value, 0))
}

// prettyPrintValue recursively formats the value.
func prettyPrintValue(value reflect.Value, depth int) string {
	switch value.Kind() {
	case reflect.Map:
		var builder strings.Builder
		keys := value.MapKeys()
		for _, key := range keys {
			builder.WriteString(strings.Repeat("  ", depth))
			builder.WriteString(fmt.Sprintf("%v: ", key.Interface()))
			builder.WriteString(prettyPrintValue(value.MapIndex(key), depth+1) + "\n")
		}
		return strings.TrimSuffix(builder.String(), "\n")
	case reflect.Slice, reflect.Array:
		var builder strings.Builder
		for i := 0; i < value.Len(); i++ {
			builder.WriteString(strings.Repeat("  ", depth))
			builder.WriteString(fmt.Sprintf("[%d]: ", i))
			builder.WriteString(prettyPrintValue(value.Index(i), depth+1) + "\n")
		}
		return strings.TrimSuffix(builder.String(), "\n")
	case reflect.Struct:
		var builder strings.Builder
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			builder.WriteString(strings.Repeat("  ", depth))
			builder.WriteString(fmt.Sprintf("%s: ", field.Name))
			builder.WriteString(prettyPrintValue(value.Field(i), depth+1) + "\n")
		}
		return strings.TrimSuffix(builder.String(), "\n")
	default:
		return fmt.Sprintf("%v", value.Interface())
	}
}
