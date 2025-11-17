package main

import (
	"demo/app-1/pkg/inet"
)

func main() {
	var url = "https://fakeapi.extendsclass.com/books/23"
	inet.CurlPrint(url)
}
