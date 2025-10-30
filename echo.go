package main

import (
	"fmt"
	"os"
)

func echo() {
	result, sep := "", " "
	for _, arg := range os.Args[1:] {
		result += sep + arg
		sep = ""
	}
	fmt.Println(result)
}