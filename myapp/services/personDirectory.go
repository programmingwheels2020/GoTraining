package services

import (
	"fmt"
)

func PrintPersonDetails() {

	fmt.Println("Database connection Opened")
	defer func() {
		fmt.Println("Data base connection closed")

	}()

	defer func() {
		fmt.Println(recover())
	}()
	panic("Error occured")

	fmt.Println("Person Details are read")
}
