package main

import (
	"fmt"
	"math"
)

func genPrimeNumbers() chan int {
	res := make(chan int)

	go func() {
		res <- 2
		for i := 3; i < 1000000; i++ {
			l := int(math.Sqrt((float64(i))))
			found := false

			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				res <- i
			}
		}
		close(res)
	}()
	return res
}

func main() {
	resCan := genPrimeNumbers()

	for n := range resCan {
		fmt.Println(n)
	}
}
