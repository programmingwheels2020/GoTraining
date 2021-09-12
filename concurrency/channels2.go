package main

import "fmt"

func recieveMessage(myChan chan string) {
	fmt.Println(<-myChan)
}

func main() {
	messages := make(chan string, 2)

	go func() {
		fmt.Println("From Go routine")
		messages <- "ONE"
		messages <- "TWO"
		messages <- "THREE"

	}()

	recieveMessage(messages)
	//recieveMessage(messages)
	//recieveMessage(messages)
	fmt.Println("From Main thread")
}
