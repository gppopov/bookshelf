package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//
// Handlers

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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	user := new(User)
	r.ParseForm()
	for k, v := range r.Form {
		if k == "username" {
			user.Username = strings.Join(v, "")
		}
		if k == "pass" {
			user.Password = strings.Join(v, "")
		}
	}
}

//
// Helpers

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

type User struct {
	Password string
	Username string
}
