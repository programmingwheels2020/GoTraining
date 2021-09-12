package main

import "fmt"

func main() {
	//var iMap map[string]int

	//iMap := map[string]int{}
	iMap := map[string]int{
		"USA": 30,
		"CHN": 14,
		"JPN": 5,
		"GER": 4,
		"IND": 3,
	}

	mMap := make(map[string]int, 10)

	fmt.Println(iMap["USA"])
	mMap["IND"] = 130
	fmt.Println(mMap)
	delete(iMap, "USA")
	v, ok := iMap["USA"]
	fmt.Println(v, ok)
	v1, ok1 := iMap["ARG"]
	fmt.Println(v1, ok1)

}
