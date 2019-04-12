package main

import (
	"fmt"

	easyjson "github.com/mailru/easyjson"
)

func main() {
	target := User{FirstName: "bar", LastName: "foo"}

	jsonBytes, err := easyjson.Marshal(&target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
