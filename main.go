package main

import (
	"demo/app-1/pkg/inet"
	"fmt"
	"os"
)

func main() {
	ch := make(chan string)

	var url1 = "https://fakeapi.extendsclass.com/books/23"
	var t1, err1 = inet.FetchDuration(url1, ch)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "%v", err1)
	}
	fmt.Printf("Duration: %d ms\n", t1)
	
	var url2 = "fakeapi.extendsclass.com/books/10"
	var t2, err2 = inet.FetchDuration(url2, ch)
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "%v", err2)
	}
	fmt.Printf("Duration: %d ms\n", t2)
}
