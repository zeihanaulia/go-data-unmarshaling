package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Target struct {
	Foo string `yaml:"foo,omitempty"`
	Bar string `yaml:"bar,omitempty"`
}

func main() {

	f, _ := os.Create("file.xml")

	target := Target{Foo: "bar1", Bar: "foo1"}
	dec := xml.NewEncoder(f)
	err := dec.Encode(&target)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(target)
}
