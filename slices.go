package main

import "fmt"

func main() {
	//var s1 = [5]int{1, 2, 3, 4, 5}
	s1 := make([]int, 5)
	s1[0] = 201
	s1[1] = 202
	s1[2] = 203
	s1[3] = 204
	/*fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s1 = append(s1, 20)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))*/
	s1 = append(s1, 21)
	s1 = append(s1, 22)
	s1 = append(s1, 23)
	s1 = append(s1, 24)
	s1 = append(s1, 25)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1[1:5])
}
