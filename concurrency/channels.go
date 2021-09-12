package main

import "fmt"

func main() {
	myChannel := make(chan string) //unbuffered
	channeltwo := make(chan string)
	go func() {
		fmt.Println("This is from Go routines")
		myChannel <- "Ping"
	}()

	go func() {
		fmt.Println("This is from Second Go routines")
		fmt.Println(<-myChannel)
		channeltwo <- "FROM SECOND"
	}()
	fmt.Println("This is from main function")
	fmt.Println(<-channeltwo)

	//fmt.Println(msg)

}
