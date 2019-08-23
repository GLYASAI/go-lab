package main

import "fmt"

type foobar struct {
	ID string
}

func main() {
	var f foobar
	f.ID = "foobar"
	fmt.Println(f)
}
