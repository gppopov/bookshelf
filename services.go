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
		fmt.Println(r.Form) // TODO: Remove after debug.
		book := new(Book)
		for k, v := range r.Form {
			if k == "name" {
				book.Name = strings.Join(v, "")
			}
			if k == "author" {
				book.Author = strings.Join(v, "")
			}
			if k == "picture" {
				book.Picurl = strings.Join(v, "")
			}
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

	if !loginUser(user) {
		// TODO: Return validation message.
	}

	// TODO: Create session for current user.
	// TODO: Create user and persist info.
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
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

// Validate password according to NIST standards
// (https://pages.nist.gov/800-63-3/sp800-63b.html#memorized-secret-verifiers)
func validatePass(user *User) bool {
	// TODO: Check if user already exists.

	// TODO: Validate password.
	// 1.) - Min 8 char, max 64 chars;
	// 2.) - Forbid space char.
	// 3.) - Check against common passwords,
	//		 dictionary words, derivatives of the username
	return true
}

func loginUser(user *User) bool {
	return true
}

// User holds system user's info.
type User struct {
	Username        string
	Password        string
	ConfirmPassword string
}
