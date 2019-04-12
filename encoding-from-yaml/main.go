package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Target struct {
	Foo string `yaml:"foo,omitempty"`
	Bar string `yaml:"bar,omitempty"`
}

func main() {

	f, _ := os.Create("file.yaml")

	target := Target{Foo: "bar1", Bar: "foo1"}
	dec := json.NewEncoder(f)
	err := dec.Encode(&target)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(target)
}
