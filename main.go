package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Quick and dirty demo to check DO apps

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server started on port 8080")

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the index!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about page!")
}
