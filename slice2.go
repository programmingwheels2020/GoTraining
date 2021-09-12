package main

import "fmt"

func main() {
	s1 := make([]int, 1024)
	//var s1 []int
	fmt.Println(cap(s1))
	//var s1 = []int{}
	s1 = append(s1, 100)
	fmt.Println(cap(s1)) //1024 - doubling  25%
	//fmt.Println(s1)
}
