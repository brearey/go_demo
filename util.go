package main

import (
	"fmt"
	"os"
	"time"
)

func PrintError(moduleName string, err error) {
	fmt.Fprintf(os.Stderr, "%s | %s error: %s\n", time.Now().Format("2006-01-02 15:04:05"), moduleName, err)
}
