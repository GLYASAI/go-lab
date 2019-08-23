package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("abc.go")
	defer file.Close()
	if err != nil {
		fmt.Println("error opening file: ", err)
		return
	}
}
