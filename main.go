package main

import (
	"demo/app-1/pkg/inet"
	"fmt"
)

func main() {
	ch1 := make(chan int)

	go inet.FetchDuration("https://fakeapi.extendsclass.com/books/1", ch1)
	go inet.FetchDuration("https://fakeapi.extendsclass.com/books", ch1)
	go inet.FetchDuration("https://fakeapi.extendsclass.com/books/3", ch1)

	fmt.Printf("first <-ch1 %v\n", <-ch1)
	fmt.Printf("second <-ch1 %v\n", <-ch1)
	fmt.Printf("third <-ch1 %v\n", <-ch1)
}
