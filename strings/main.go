package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("/foo/bar/", "/dummy"))
}
