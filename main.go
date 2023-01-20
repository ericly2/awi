package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = 5000

var SERVER_URL = fmt.Sprintf("localhost:%d", PORT)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running on %s", SERVER_URL)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	http.HandleFunc("/book", GetBook)
	fmt.Printf("Starting the server on %s\n", SERVER_URL)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
