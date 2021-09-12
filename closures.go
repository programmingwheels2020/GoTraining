package main

import "fmt"

func functionone() func() int {
	return func() int {
		fmt.Println("Sub function executed")
		return 10
	}
}

func main() {
	subfunc := functionone()
	s := subfunc()
	fmt.Println(s)
}
