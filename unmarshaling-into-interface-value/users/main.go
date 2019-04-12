package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Target struct {
	Username string                 `json:"username,omitempty"`
	Height   float64                `json:"height,omitempty"`
	Address  map[string]interface{} `json:"address,omitempty"`
	Interest []interface{}          `json:"interest,omitempty"`
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
