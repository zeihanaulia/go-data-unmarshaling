package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Target struct {
	Foo string `json:"foo,omitempty"`
	Bar string `json:"bar,omitempty"`
}

func main() {

	f, err := os.Open("file.json")
	if err != nil {
		panic(err)
	}

	target := Target{}
	dec := json.NewDecoder(f)

	err = dec.Decode(&target)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(target)
}
