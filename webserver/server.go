package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", FirstHandler)
	http.ListenAndServe(":8000", nil)

}

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
