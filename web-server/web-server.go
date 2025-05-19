package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handle the root URL ("/") with a simple "Hello, world!" message.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r.Method, r.URL)
		if r.URL.Path == "/" {
			fmt.Println("Hello, world!")
			w.Write([]byte("Hello, world!"))
		}
	})

	// Start the web server on port 8080.
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
