package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	type ToDo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)

	/*decoder := json.NewDecoder(resp.Body)
	var t ToDo
	err = decoder.Decode(&t)
	if err != nil {
		panic(err)
	}*/

	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))

	var t ToDo

	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	log.Println(t)

	/*
		scanner := bufio.NewScanner(resp.Body)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}*/

}
