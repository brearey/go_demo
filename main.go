package main

import (
	"demo/app-1/pkg/password"
	"fmt"
)

func main() {
	var password, err = password.Generate(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(password)
	}
}
