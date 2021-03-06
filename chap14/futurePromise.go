package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(path string) chan string {
	promise := make(chan string)
	go func() {
		content, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Println("Error!")
			close(promise)
		} else {
			promise <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSrc chan string) chan []string {
	promise := make(chan []string)
	go func() {
		var result []string

		for _, line := range strings.Split(<-futureSrc, "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise

}

func main() {
	futureSrc := readFile("futurePromise.go")
	futureFuncs := printFunc(futureSrc)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}
