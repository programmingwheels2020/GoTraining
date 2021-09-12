package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		fmt.Println("Go routine One")
		ch1 <- "ONE"
	}()

	go func() {

		fmt.Println("Go routine Two")
		ch2 <- "TWO"
	}()

	for i := 0; i < 2; i++ {

		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)

		}
	}
}
