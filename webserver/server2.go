package main

import (
	"fmt"
	"net/http"
)

type HelloWeb int

func (h *HelloWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the site")
}

var hello HelloWeb

func main() {
	http.ListenAndServe(":8000", &hello)
}
