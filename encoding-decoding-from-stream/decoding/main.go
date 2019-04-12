package main

import (
	"encoding/json"
	"fmt"
)

type Target struct {
	Foo string `json:"foo,omitempty"`
}

func main() {

	jsonBytes := []byte(`{"foo":"bar"}`)
	target := Target{}

	err := json.Unmarshal(jsonBytes, &target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
