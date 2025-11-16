package main

import "fmt"

func main() {
	var password, err = Generate(10)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(password)
	}
}
