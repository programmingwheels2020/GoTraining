package main

import "fmt"

func square(y *int) {

	fmt.Println(*y)
	*y = *y * *y
	fmt.Println(*y)
}

func squareValue(y int) {
	y = y * y
	fmt.Println(y)
}

func main() {
	var x = 10
	//xPtr := &x
	//var xPtr *int
	//xPtr = &x
	//fmt.Println(x)
	//fmt.Println(xPtr)
	//square(&x)
	//fmt.Println(x)
	squareValue(x)
	fmt.Println(x)
}
