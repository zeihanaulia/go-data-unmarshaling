package main

import (
	"fmt"
	"os"

	"github.com/vmihailenco/msgpack"
)

type Target struct {
	Foo string `msgpack:"foo,omitempty"`
	Bar string `msgpack:"bar,omitempty"`
}

func main() {

	f, _ := os.Create("file.bin")

	target := Target{Foo: "bar1", Bar: "foo1"}
	dec := msgpack.NewEncoder(f)
	err := dec.Encode(&target)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(target)
}
