package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
