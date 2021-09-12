package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) (error, int) {
	if a < 0 || b < 0 {
		return errors.New("Negative Number"), 0
	}
	sum := a + b
	return nil, sum
	//fmt.Println(sum)
}

func calculate(a, b int) (int, int) {
	return a + b, a - b
}

func sum(nums ...int) int {
	fmt.Println(nums)
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func main() {
	//fmt.Println("Inside Main")
	/*err, sum := add(-2, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}*/

	//sum, subst := calculate(10, 5)
	//fmt.Println(sum, subst)
	nums := []int{8, 9, 0}
	total := sum(nums...)
	//total := sum(3, 4, 5, 6, 8)
	fmt.Println(total)
}
