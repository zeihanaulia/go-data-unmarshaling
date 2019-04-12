package main

import (
	"encoding/json"
	"fmt"
)

type Target struct {
	Foo string `json:"foo,omitempty"`
}

func main() {
	target := Target{Foo: "bar"}

	jsonBytes, err := json.Marshal(&target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
