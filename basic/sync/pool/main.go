package main

import (
	"fmt"
	"sync"
)

func main() {
	newFunc := func() interface{} {
		return 0
	}
	pool := sync.Pool{New: newFunc}

	pool.Put(10)
	pool.Put(11)
	pool.Put(12)
	for i := 0; i < 5; i++ {
		v := pool.Get()
		fmt.Printf("value: %v\n", v)
		/**
		value: 10
		value: 12
		value: 11
		value: 0 Put
		value: 0
		 */
	}
}
