package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("Created:%d\n", count)
		},
	}

	pool.Put("Manualy added:1")
	pool.Put("Manualy added:2")

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())

}
