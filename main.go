package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	// set routes
	http.HandleFunc("/book", addBookHandler)
	http.HandleFunc("/books", listBooksHandler)
	http.HandleFunc("/login", loginHandler)
	err := http.ListenAndServeTLS(":8099", "cert/localhost.pem", "cert/localhost.key", nil)
	logErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func logErr(err error) {
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Book type holds books info.
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Picurl string `json:"picurl"`
}
