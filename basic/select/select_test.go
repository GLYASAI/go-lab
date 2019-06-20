package sel

import (
	"context"
	"testing"
)

func TestFoobar(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	ich := make(chan int)
	go Foobar(ctx, ich)
	for i:=-5; i < 5; i++ {
		ich <- i
	}
	cancelFunc()
}
