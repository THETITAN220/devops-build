package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	fmt.Println("Server running on http://localhost:3030")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Println("Error starting the server:", err)
	}
}
