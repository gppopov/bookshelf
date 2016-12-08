package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func listBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := generateBooks(20)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// TODO: Get book.
		// TODO: Return book info as json.
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		// TOOD: Parse form and validate inputs.
		// TOOD: Persist data.
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}
