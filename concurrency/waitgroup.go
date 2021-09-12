package main

import (
	"fmt"
	"sync"
)

func message(msg string) {
	for i := 0; i < 20; i++ {
		fmt.Println(msg, ":", i)
	}
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Execution Starts")

	wg.Add(1)
	go func() {
		message("ONE")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		message("TWO")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Finished")
}
