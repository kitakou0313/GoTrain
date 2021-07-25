package main

import "fmt"

func genOddNum(limit int) chan int {
	resChan := make(chan int)

	go func() {
		for i := 1; i <= limit; i += 2 {
			resChan <- i
		}
	}()

	return resChan
}

func genEvenNum(limit int) chan int {
	resChan := make(chan int)

	go func() {
		for i := 0; i <= limit; i += 2 {
			resChan <- i
		}
	}()

	return resChan
}

func main() {

	oddChan := genOddNum(100)
	evenChan := genEvenNum(100)

	for {
		select {
		case res := <-oddChan:
			fmt.Println("Odd num:", res)
		case res := <-evenChan:
			fmt.Println("Even num:", res)
		default:
			break
		}
	}

}
