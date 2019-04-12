package main

import (
	"encoding/json"
	"fmt"
)

type Target struct {
	Foo interface{} `json:"foo,omitempty"`
}

func main() {

	jsonBytes := []byte(`[{"foo":"bar"}, {"foo": 123}]`)
	target := []Target{}

	err := json.Unmarshal(jsonBytes, &target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", target)

	// Penting sebelum menggunakan interface{} musti kenal dulu
	// https://golang.org/ref/spec#Type_assertions

	for _, t := range target {
		switch v := t.Foo.(type) {
		case int:
			fmt.Printf("\n foo is int %d", v)
		case float64:
			fmt.Printf("\nfoo is float64 %T", v)
		case string:
			fmt.Printf("\nfoo is string: %s\n", v)
		default:
			fmt.Printf("foo is something else: %T\n", v)
		}
	}

}
