package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func about(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		fmt.Fprintf(w, "About")
	case http.MethodPost:
		fmt.Fprintf(w, "About")
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/about", about)

	http.ListenAndServe(":8080", mux)
}
