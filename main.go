package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "localhost:5000 up and running")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := Book{
		Title:  "The Gunslinger",
		Author: "Stephen King",
		Pages:  304,
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)
	fmt.Println("Starting the server on localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
