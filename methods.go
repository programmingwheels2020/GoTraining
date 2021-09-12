package main

import "fmt"

type Student struct {
	name string
	age  int
	mark map[string]float64
}

func (s Student) CalculateTotal() float64 {
	sum := 0.0

	for _, v := range s.mark {
		//fmt.Println(k, v)
		sum += v
	}
	return sum
}

func (s *Student) UpdateAge() {
	fmt.Println(s)
	s.age = 37
}

func main() {
	s1 := Student{
		"Lionell Messi", 17, map[string]float64{
			"PHY": 99.0,
			"CHE": 79.0,
			"BIO": 87.0,
		},
	}

	fmt.Println(s1)
	s1.age = 37
	fmt.Println(s1)

	/*total := s1.CalculateTotal()
	fmt.Println("Total Mark is ", total)
	s1.UpdateAge()
	fmt.Println(s1)*/
}
