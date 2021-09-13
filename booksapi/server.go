package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Language string `json:"language"`
}

var books []Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	fmt.Fprintf(w, `{"message":"Successfully created book"}`)
}

func GetAllBooksByLanguage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var booksByLang []Book

	fmt.Println(params["language"])
	for _, v := range books {
		if v.Language == params["language"] {
			booksByLang = append(booksByLang, v)
		}
	}
	json.NewEncoder(w).Encode(booksByLang)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{language}", GetAllBooksByLanguage).Methods("GET")
	router.HandleFunc("/books", CreateBooks).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
