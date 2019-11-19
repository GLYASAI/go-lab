package noblock

import "fmt"

// NoBlock -
func NoBlock(listeners map[string]chan<- int) {
	for {
		for _, pech := range listeners {
			select {
			case pech <- 1:
				fmt.Print("hihi\n")
			default:
				fmt.Print("default\n")
			}
		}
		fmt.Print("okokok\n")
	}
}
