package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Target struct {
	Foo string `yaml:"foo,omitempty"`
	Bar string `yaml:"bar,omitempty"`
}

func main() {

	f, _ := os.Create("file.yaml")

	target := Target{Foo: "bar1", Bar: "foo1"}
	dec := yaml.NewEncoder(f)
	err := dec.Encode(&target)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(target)
}
