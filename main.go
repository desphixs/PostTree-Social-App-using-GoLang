package main

import (
	"fmt"
	"net/http"
)

func homeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to social go")
}

func main () {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Starting server at :8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server at :8080")
	}
}