package main

import "fmt"

func main() {
	var arr = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, num := range arr {
		fmt.Println(num)
		fmt.Println(i)
	}

	fruits := map[string]float64{
		"Apple":  200.50,
		"Orange": 100.90,
		"Grapes": 400.40,
	}

	for k, v := range fruits {
		fmt.Println(k)
		fmt.Println(v)
	}

	for i, s := range "Hello World" {
		fmt.Println(i)
		fmt.Println(s)
	}
}
