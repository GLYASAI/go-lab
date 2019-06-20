package sel

import (
	"context"
	"fmt"
)

// Foobar -
func Foobar(ctx context.Context, ich chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-ich:
			if event == -1 {
				fmt.Printf("Received %d, continue.\n", event)
				// continue
				break
			}
			fmt.Printf("Received %d.\n", event)
		}
	}
}
