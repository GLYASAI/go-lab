package main

import (
	"fmt"
	"os"
)

func main() {
	logFile, err := os.Stat("../README.md")
	if err != nil {
		fmt.Printf("read file: %v/n", err)
		return
	}
	fmt.Println(logFile.Name())
}
