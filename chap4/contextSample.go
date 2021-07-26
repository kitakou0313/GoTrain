package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("Start sub.")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("Sub() is finished")
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
