package main

import "fmt"

func main() {
	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	i := 0
	/*for i < 10 {
		fmt.Println(i)
		i++
	}*/

	for {

		i++
		if i%2 == 0 {
			continue
			//fmt.Println("Hi")
		} else {
			fmt.Println("Hello")
		}
		if i > 100 {
			break
		}

	}

}
