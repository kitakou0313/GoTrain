package main

import (
	"fmt"
)

func main() {
	tasks := make(chan string)

	go func() {
		tasks <- "Send in task"
	}()

	res, isOpen := <-tasks
	fmt.Println(res, isOpen)
}
