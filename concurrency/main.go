package main

import (
	"fmt"
	"time"
)

func funcone(msg string) {
	for i := 0; i < 50; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	/*fmt.Println("This is One")
	go func() {
		fmt.Println("This is Two")
	}()

	fmt.Println("This is Three")
	time.Sleep(8 * time.Second) */

	//funcone("Direct")
	go funcone("goroutine")
	go funcone("AnotherGoroutine")

	time.Sleep(2 * time.Second)
	fmt.Println("Finished")

}
