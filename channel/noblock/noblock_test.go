package noblock

import (
	"fmt"
	"testing"
)

func TestNoblock(t *testing.T) {
	ch := make(chan int, 1)
	listeners := make(map[string]chan<- int)
	go func() {
		for {
			select {
			case c := <-ch:
				fmt.Printf("received: %d\n", c)
			default:
				fmt.Print("received nothing\n")
			}
		}
	}()
	listeners["foobar"] = ch
	NoBlock(listeners)

}
