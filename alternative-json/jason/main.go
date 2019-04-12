package main

import (
	"fmt"

	"github.com/antonholmquist/jason"
)

func main() {
	jsonBytes := []byte(`{
		"firstName":"zei",
		"lastName": "han",
		"email":"zei@han.com",
		"address": {
			"street": "jalan",
			"city": {
				"name":"jakarta",
				"code":"14045"
			}
		}
	}`)

	user, err := jason.NewObjectFromBytes(jsonBytes)
	if err != nil {
		fmt.Println(err)
	}

	firstName, _ := user.GetString("firstName")
	fmt.Println(firstName)

	cityName, _ := user.GetString("address", "city", "name")
	fmt.Println(cityName)
}
