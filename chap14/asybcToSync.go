package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	inputs := make(chan []byte, 2)

	wg.Add(2)

	go func() {
		a, _ := ioutil.ReadFile("a.txt")
		inputs <- a
		wg.Done()
	}()

	go func() {
		b, _ := ioutil.ReadFile("b.txt")
		inputs <- b
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(<-inputs)
	fmt.Println(<-inputs)

}
