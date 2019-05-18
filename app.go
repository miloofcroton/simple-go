package main

import (
	"fmt"
	"net/http"
	"os"

	// run: `go get github.com/miloofcroton/simple-go/api`
	"github.com/miloofcroton/simple-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", api.HelloHandleFunc)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	// adding localhost to avoid getting the MacOS firewall notification
	return "localhost:" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is the Book API")
}
