package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var counter int

func increment(msg string) {
	for i := 0; i < 30; i++ {
		mutex.Lock()
		x := counter
		x++
		time.Sleep(2 * time.Second)
		counter = x
		fmt.Println(msg, ":", counter)
		mutex.Unlock()
	}
}

func main() {
	fmt.Println("Execution Started")
	wg.Add(2)
	go func() {
		increment("ONE")
		wg.Done()
	}()

	go func() {
		increment("TWO")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Finished")

}
