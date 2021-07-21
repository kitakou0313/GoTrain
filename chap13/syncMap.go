package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}

	smap.Store(1, 2)
	smap.Store("Hello", "World")

	res, ok := smap.Load(1)
	if !ok {
		fmt.Println(ok)
	}
	fmt.Println(res)

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	smap.LoadOrStore(1, 2)
	smap.LoadOrStore(1, 3)

	res, ok = smap.Load(1)
	if !ok {
		fmt.Println(ok)
	}
	fmt.Println(res)

}
