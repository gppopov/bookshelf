package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/book", addBookHandler) // set router
	http.HandleFunc("/books", listBooksHandler)
	//http.HandleFunc("/AddBook", addBookHandler)
	//http.ListenAndServe(":8099", http.FileServer(http.Dir("public")))
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func generateBooks(length int) []Book {
	books := make([]Book, length, length)
	for i := 0; i < length; i++ {
		books[i] = Book{
			ID:     i,
			Name:   fmt.Sprint("name-", i),
			Author: fmt.Sprint("author-", i),
			Picurl: fmt.Sprint("pic-", i),
		}
	}

	return books
}

// Book type holds books info.
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Picurl string `json:"picurl"`
}
