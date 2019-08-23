package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		timer := time.NewTimer(1 * time.Second)
		for {
			ch <- "hahaha"
			timer.Reset(1 * time.Second)
		}
	}()

	for item := range <-ch {
		fmt.Println(item)
	}

	
}
